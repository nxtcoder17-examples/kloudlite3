import { createContext, useContext, useEffect, useMemo } from 'react';
import { ImmerHook, useImmer } from 'use-immer';
import { ChildrenProps } from '~/components/types';
import { NonNullableString } from '~/root/lib/types/common';
import {
  AppIn,
  Github__Com___Kloudlite___Operator___Apis___Crds___V1__AppContainerIn as AppSpecContainersIn,
  BuildIn,
} from '~/root/src/generated/gql/server';
import { mapper } from '~/components/utils';
import { parseNodes } from '~/console/server/r-utils/common';
import { IApp } from '../server/gql/queries/app-queries';

const defaultApp: AppIn & { build?: BuildIn } = {
  metadata: {
    name: '',
    annotations: [],
  },
  spec: {
    containers: [
      {
        image: '',
        name: 'container-1',
        env: [],
      },
    ],
  },
  displayName: '',
};

const defaultBuild: BuildIn = {
  name: '',
  source: {
    branch: '',
    provider: 'github',
    repository: '',
  },
  buildClusterName: '',
  spec: {
    registry: {
      repo: {
        name: '',
        tags: [],
      },
    },
    resource: {
      cpu: 500,
      memoryInMb: 1000,
    },
  },
};

export type ISetState<T = any> = (fn: ((val: T) => T) | T) => void;
type ISetContainer<T = any> = (fn: ((val: T) => T) | T, index?: number) => void;

const CreateAppContext = createContext<any>(null);

export type createAppEnvPage =
  | 'environment_variables'
  | 'config_mounts'
  | NonNullableString;

interface IappState {
  completePages: {
    [key: string]: boolean;
  };
  activeContIndex: number;
  envPage: createAppEnvPage;
  page: number;
  app: AppIn;
  buildData?: BuildIn | null | undefined;
  readOnlyApp?: IApp;
}

export const useAppState = () => {
  const [state, setState] = useContext<ImmerHook<IappState>>(CreateAppContext);

  const {
    app,
    page,
    envPage,
    activeContIndex,
    completePages,
    buildData,
    readOnlyApp,
  } = state;

  const getContainer = (index: number = activeContIndex) => {
    if (!index) {
      // eslint-disable-next-line no-param-reassign
      index = 0;
    }
    return (
      app.spec.containers[index] || {
        name: `container-${index}`,
        image: '',
      }
    );
  };

  const setReadOnlyApp: ISetState<IApp> = (fn) => {
    if (typeof fn === 'function') {
      setState((s) => ({ ...s, readOnlyApp: fn(s.readOnlyApp) }));
    } else {
      setState((s) => ({ ...s, readOnlyApp: fn }));
    }
  };

  const setApp: ISetState<typeof app> = (fn) => {
    if (typeof fn === 'function') {
      setState((s) => ({ ...s, app: fn(s.app) }));
    } else {
      setState((s) => ({ ...s, app: fn }));
    }
  };

  const setContainer: ISetContainer<AppSpecContainersIn> = (
    fn,
    index = activeContIndex
  ) => {
    const containers = [...(app.spec.containers || [])];

    if (typeof fn === 'function') {
      const container = containers[index];
      containers[index] = fn(
        container || {
          name: `container-${index}`,
          image: '',
        }
      );
    } else {
      containers[index] = fn;
    }

    setApp((a) => {
      const app = {
        ...a,
        spec: {
          ...a.spec,
          containers,
        },
      };
      return app;
    });
  };

  const setBuildData: ISetState<BuildIn> | null = (fn) => {
    if (typeof fn === 'function') {
      // @ts-ignore
      setState((s) => ({ ...s, buildData: fn(s.buildData) }));
    } else {
      setState((s) => ({ ...s, buildData: fn }));
    }
  };

  const setPage: ISetState<number> = (fn) => {
    if (typeof fn === 'function') {
      setState((s) => ({ ...s, page: fn(s.page) }));
    } else {
      setState((s) => ({ ...s, page: fn }));
    }
  };

  const setEnvPage: ISetState<createAppEnvPage> = (fn) => {
    if (typeof fn === 'function') {
      setState((s) => ({ ...s, envPage: fn(s.envPage) }));
    } else {
      setState((s) => ({ ...s, envPage: fn }));
    }
  };

  useEffect(() => {
    if (!page) {
      setPage(1);
    }
    if (!envPage) {
      setEnvPage('environment_variables');
    }

    if (!app) {
      setApp(defaultApp);
    }
    // if (!completePages) {
    // }

    if (!activeContIndex || activeContIndex !== 0) {
      setState((s) => ({
        ...s,
        activeContIndex: 0,
      }));
    }
  }, []);

  const setServices: ISetState<NonNullable<AppIn['spec']['services']>> = (
    fn
  ) => {
    if (typeof fn === 'function') {
      setApp((a) => ({
        ...a,
        spec: {
          ...a.spec,
          services: fn(a.spec.services || []),
        },
      }));
    } else {
      setApp((a) => ({
        ...a,
        spec: {
          ...a.spec,
          services: fn,
        },
      }));
    }
  };

  const isPageComplete = (page: number) => {
    if (completePages) return completePages[page];

    setState((s) => {
      return {
        ...s,
        completePages: {},
      };
    });
    return false;
  };

  const markPageAsCompleted = (page: number) => {
    setState((s) => {
      return {
        ...s,
        completePages: {
          ...s.completePages,
          [page]: true,
        },
      };
    });
  };

  const resetState = (iApp?: AppIn) => {
    setState({
      page: 1,
      app: iApp || defaultApp,
      completePages: {},
      envPage: 'environment_variables',
      activeContIndex: 0,
      buildData: defaultBuild,
    });
  };

  const resetBuildData = () => {
    setBuildData(defaultBuild);
  };

  type IparseNodes = {
    edges: Array<{ node: any }>;
  };

  const getRepoMapper = (resources: IparseNodes | undefined) => {
    return mapper(parseNodes(resources), (val) => ({
      label: val.name,
      value: val.name,
      accName: val.accountName,
    }));
  };

  const getRepoName = (imageUrl: string) => {
    const parts: string[] = imageUrl.split(':');
    const repoParts: string[] = parts[0].split('/');
    if (repoParts.length === 1) {
      return repoParts[repoParts.length - 1];
    }
    const repoSlicePart: string[] = repoParts.slice(2);
    return repoSlicePart.join('/');
  };

  const getImageTag = (imageUrl: string) => {
    const parts: string[] = imageUrl.split(':');
    // logger.log('image tag', parts[1]);
    return parts[1];
  };

  return {
    resetState,
    completePages,
    isPageComplete,
    markPageAsCompleted,
    app: app || defaultApp,
    setApp,
    envPage,
    setEnvPage,
    page,
    setPage,
    state,
    setState,
    getContainer,
    setContainer,
    activeContIndex: activeContIndex || 0,
    services: app.spec.services || [],
    setServices,
    getRepoMapper,
    getRepoName,
    getImageTag,
    setBuildData,
    buildData,
    resetBuildData,
    readOnlyApp,
    setReadOnlyApp,
  };
};

export const clearAppState = () => {
  if (typeof window === 'undefined') return;
  sessionStorage.removeItem('state');
};

export const AppContextProvider = ({
  children,
  initialAppState,
}: ChildrenProps & { initialAppState?: AppIn & { build?: BuildIn } }) => {
  const loadSession = () => {
    if (typeof window === 'undefined')
      return {
        app: defaultApp,
        readOnlyApp: defaultApp,
        buildData: defaultApp?.build,
      };
    if (initialAppState) {
      return {
        app: initialAppState,
        readOnlyApp: initialAppState,
        buildData: initialAppState?.build,
      };
    }
    const stateString =
      sessionStorage.getItem('state') ||
      JSON.stringify({
        app: defaultApp,
        readOnlyApp: defaultApp,
      });

    try {
      const data = JSON.parse(stateString);
      return data || {};
    } catch (_) {
      return {};
    }
  };
  const [state, setState] = useImmer<IappState>(() => {
    return loadSession();
  });

  useEffect(() => {
    if (typeof window === 'undefined' || initialAppState) return;
    // logger.log(initialAppState, 'hrere');
    sessionStorage.setItem('state', JSON.stringify(state || {}));
  }, [state]);

  return (
    <CreateAppContext.Provider
      value={useMemo(() => [state, setState], [state, setState])}
    >
      {children}
    </CreateAppContext.Provider>
  );
};
