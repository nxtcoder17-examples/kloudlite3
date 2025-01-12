import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useRef,
  useState,
} from 'react';
import { ChildrenProps } from '@kloudlite/design-system/types';
import ReconnectingWebSocket from 'reconnecting-websocket';
import logger from '~/root/lib/client/helpers/log';
import useDebounce from '~/root/lib/client/hooks/use-debounce';
import { socketUrl } from '~/root/lib/configs/base-url.cjs';
import { sleep } from '~/root/lib/utils/common';
import { usePulsableLoading } from '../../components/pulsable';

type IFor = 'logs' | 'resource-update';

export interface ISocketResp<T = any> {
  type: 'response' | 'error' | 'info';
  for: IFor;
  message: string;
  id: string;
  data: T;
}

type IData = {
  event?: 'subscribe' | 'unsubscribe';
  id: string;
};

interface ISocketMsg<T extends IData> {
  for: IFor;
  data: T;
}

interface IResponses {
  [key: string | IFor]: {
    [id: string]: ISocketResp[];
  };
}

type IsendMsg = <T extends IData>(msg: ISocketMsg<T>) => void;

const Context = createContext<{
  responses: IResponses;
  errors: IResponses;
  infos: IResponses;
  sendMsg: IsendMsg;
  clear: (msg: ISocketMsg<IData>) => void;
}>({
  clear: () => {},
  responses: {},
  errors: {},
  infos: {},
  sendMsg: () => {},
});

export const useSubscribe = <T extends IData>(
  msg: ISocketMsg<T> | ISocketMsg<T>[],
  dep: never[]
) => {
  const {
    sendMsg,
    responses,
    infos: mInfos,
    errors: mErrors,
    clear,
  } = useContext(Context);

  const isPulsableLoading = usePulsableLoading();

  const [resp, setResp] = useState<ISocketResp[]>([]);
  const [subscribed, setSubscribed] = useState(false);
  const [errors, setErrors] = useState<ISocketResp[]>([]);
  const [infos, setInfos] = useState<ISocketResp[]>([]);

  useEffect(() => {
    (async () => {
      if (Array.isArray(msg)) {
        setResp(resp);

        const tr: ISocketResp[] = [];
        const terr: ISocketResp[] = [];
        const ti: ISocketResp[] = [];

        for (let k = 0; k < msg.length; k += 1) {
          const m = msg[k];

          tr.push(...(responses[m.for]?.[m.data.id || 'default'] || []));
          terr.push(...(mErrors[m.for]?.[m.data.id || 'default'] || []));
          ti.push(...(mInfos[m.for]?.[m.data.id || 'default'] || []));
        }

        setResp(tr);
        setErrors(terr);
        setInfos(ti);

        if (tr.length || ti.length) {
          setSubscribed(true);
        }
        return;
      }
      const tempResp = responses[msg.for]?.[msg.data.id || 'default'] || [];
      setResp(tempResp);

      setErrors(mErrors[msg.for]?.[msg.data.id || 'default'] || []);

      const tempInfo = mInfos[msg.for]?.[msg.data.id || 'default'] || [];
      setInfos(tempInfo);

      if (tempResp.length || tempInfo.length) {
        setSubscribed(true);
      }
    })();
  }, [responses, mInfos, mErrors]);

  useDebounce(
    () => {
      if (isPulsableLoading) {
        return () => {};
      }

      logger.log('subscribing');
      if (Array.isArray(msg)) {
        msg.forEach((m) => {
          sendMsg({ ...m, data: { ...m.data, event: 'subscribe' } });
        });
      } else {
        sendMsg({ ...msg, data: { ...msg.data, event: 'subscribe' } });
      }

      return () => {
        logger.log('unsubscribing');
        if (Array.isArray(msg)) {
          msg.forEach((m) => {
            clear(m);
            setSubscribed(false);
            sendMsg({ ...m, data: { ...m.data, event: 'unsubscribe' } });
          });
          return;
        }

        clear(msg);
        setSubscribed(false);
        sendMsg({ ...msg, data: { ...msg.data, event: 'unsubscribe' } });
      };
    },
    1000,
    [...dep]
  );

  return {
    responses: resp,
    subscribed,
    infos,
    errors,
  };
};

export const SockProvider = ({ children }: ChildrenProps) => {
  const sockPromise = useRef<Promise<ReconnectingWebSocket> | null>(null);

  const [responses, setResponses] = useState<IResponses>({});
  const [errors, setErrors] = useState<IResponses>({});
  const [infos, setInfos] = useState<IResponses>({});

  const setResp = useCallback((resp: ISocketResp) => {
    setResponses((s) => {
      const key = resp.for;
      const { id } = resp;
      return {
        ...s,
        [key]: {
          ...s[key],
          [id]: [...(s[key]?.[id] || []), resp],
        },
      };
    });
  }, []);

  const setError = useCallback((resp: ISocketResp) => {
    setErrors((s) => {
      const key = resp.for;
      const { id } = resp;
      return {
        ...s,
        [key]: {
          ...s[key],
          [id]: [...(s[key]?.[id] || []), resp],
        },
      };
    });
  }, []);

  const setInfo = useCallback((resp: ISocketResp) => {
    setInfos((s) => {
      const key = resp.for;
      const { id } = resp;
      return {
        ...s,
        [key]: {
          ...s[key],
          [id]: [...(s[key]?.[id] || []), resp],
        },
      };
    });
  }, []);

  const onMessage = useCallback((msg: any) => {
    try {
      const m: ISocketResp = JSON.parse(msg.data as string);

      switch (m.type) {
        case 'response':
          setResp(m);
          break;

        case 'info':
          setInfo(m);
          break;
        case 'error':
          logger.error('error from socket:', m);
          setError(m);
          break;
        default:
          logger.log('unknown message', m);
      }
    } catch (err) {
      console.error(err);
      logger.log('error parsing:', msg.data);
    }
  }, []);

  const getSocket = () => {
    return new Promise<ReconnectingWebSocket>((res, rej) => {
      try {
        // eslint-disable-next-line new-cap
        const w = new ReconnectingWebSocket(`${socketUrl}/ws`, '', {});

        w.onmessage = onMessage;

        w.onopen = () => {
          res(w);
        };

        w.onerror = (e) => {
          rej(e);
        };

        w.onclose = () => {
          rej();
        };
      } catch (e) {
        rej(e);
      }
    });
  };

  useDebounce(
    () => {
      if (typeof window !== 'undefined') {
        try {
          sockPromise.current = getSocket();
        } catch (e) {
          logger.error(e);
        }
      }
    },
    1000,
    []
  );

  const sendMsg = useCallback(
    async <T extends IData>(msg: ISocketMsg<T>) => {
      while (!sockPromise.current) {
        logger.log('no socket connection waiting');
        // eslint-disable-next-line no-await-in-loop
        await sleep(1000);
      }
      try {
        const w = await sockPromise.current;
        if (!w) {
          logger.log('no socket connection');
          return;
        }

        w.send(JSON.stringify(msg));
      } catch (err) {
        console.error(err);
      }
    },
    [sockPromise.current]
  );

  const clear = useCallback(
    <T extends IData>(msg: ISocketMsg<T>) => {
      setResponses((s) => {
        const key = msg.for;
        const id = msg.data.id || 'default';
        return {
          ...s,
          [key]: {
            ...s[key],
            [id]: [],
          },
        };
      });

      setErrors((s) => {
        const key = msg.for;
        const id = msg.data.id || 'default';
        return {
          ...s,
          [key]: {
            ...s[key],
            [id]: [],
          },
        };
      });

      setInfos((s) => {
        const key = msg.for;
        const id = msg.data.id || 'default';
        return {
          ...s,
          [key]: {
            ...s[key],
            [id]: [],
          },
        };
      });
    },
    [responses]
  );

  return (
    <Context.Provider
      value={useMemo(() => {
        return {
          clear,
          responses,
          errors,
          infos,
          sendMsg,
        };
      }, [responses, errors, infos, sendMsg])}
    >
      {children}
    </Context.Provider>
  );
};
