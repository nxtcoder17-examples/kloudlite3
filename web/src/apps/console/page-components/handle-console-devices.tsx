/* eslint-disable no-nested-ternary */
/* eslint-disable react/destructuring-assignment */
import {
  ArrowLineDown,
  ArrowRight,
  ChevronLeft,
  ChevronRight,
  Plus,
  SmileySad,
  X,
} from '~/console/components/icons';
import { useEffect, useState } from 'react';
import { Button, IconButton } from '@kloudlite/design-system/atoms/button';
import { NumberInput } from '@kloudlite/design-system/atoms/input';
import { usePagination } from '@kloudlite/design-system/molecule/pagination';
import Popup from '@kloudlite/design-system/molecule/popup';
import { toast } from '@kloudlite/design-system/molecule/toast';
import { cn } from '@kloudlite/design-system/utils';
import List from '~/console/components/list';
import NoResultsFound from '~/console/components/no-results-found';
import QRCode from '~/console/components/qr-code';
import { useConsoleApi } from '~/console/server/gql/api-provider';
import useForm, { dummyEvent } from '~/root/lib/client/hooks/use-form';
import Yup from '~/root/lib/server/helpers/yup';
import { downloadFile } from '~/console/utils/commons';
import CodeView from '~/console/components/code-view';
import { InfoLabel } from '~/console/components/commons';
import { parseValue } from '~/console/page-components/util';
import { NameIdView } from '~/console/components/name-id-view';
import useCustomSwr from '~/root/lib/client/hooks/use-custom-swr';
import Select from '@kloudlite/design-system/atoms/select';
import { Link } from '@remix-run/react';
import { ConsoleApiType } from '../server/gql/saved-queries';
import ExtendedFilledTab from '../components/extended-filled-tab';
import { LoadingPlaceHolder } from '../components/loading';

interface IExposedPorts {
  targetPort?: number;
  port?: number;
}

interface IExposedPortList {
  exposedPorts: IExposedPorts[];
  onDelete: (exposedPorts: IExposedPorts) => void;
}
const ExposedPortList = ({
  exposedPorts,
  onDelete = (_) => _,
}: IExposedPortList) => {
  const itemsPerPage = 4;

  const { page, hasNext, hasPrevious, onNext, onPrev, setItems } =
    usePagination({
      items: exposedPorts,
      itemsPerPage,
    });

  useEffect(() => {
    setItems(exposedPorts);
  }, [exposedPorts]);
  return (
    <div className="flex flex-col gap-lg bg-surface-basic-default">
      {exposedPorts.length > 0 && (
        <List.Root
          className="min-h-[265px] !shadow-none"
          header={
            <div className="flex flex-row items-center w-full">
              <div className="text-text-strong bodyMd flex-1">
                Exposed ports
              </div>
              <div className="flex flex-row items-center">
                <IconButton
                  icon={<ChevronLeft />}
                  size="xs"
                  variant="plain"
                  onClick={() => onPrev()}
                  disabled={!hasPrevious}
                />
                <IconButton
                  icon={<ChevronRight />}
                  size="xs"
                  variant="plain"
                  onClick={() => onNext()}
                  disabled={!hasNext}
                />
              </div>
            </div>
          }
        >
          {page.map((ep, index) => {
            return (
              <List.Row
                className={cn({
                  '!border-b': index < itemsPerPage - 1,
                  '!rounded-b-none': index < itemsPerPage - 1,
                })}
                key={ep.port}
                columns={[
                  {
                    key: `${index}-column-2`,
                    className: 'flex-1',
                    render: () => (
                      <div className="flex flex-row gap-md items-center bodyMd text-text-soft">
                        <span>Exposed: </span>
                        {ep.port}
                        <ArrowRight size={16} />
                        <span>Target: </span>
                        {ep.targetPort}
                      </div>
                    ),
                  },
                  {
                    key: `${index}-column-3`,
                    render: () => (
                      <div>
                        <IconButton
                          icon={<X />}
                          variant="plain"
                          size="sm"
                          onClick={() => {
                            onDelete(ep);
                          }}
                        />
                      </div>
                    ),
                  },
                ]}
              />
            );
          })}
        </List.Root>
      )}
      {exposedPorts.length === 0 && (
        <div className="rounded border-border-default border min-h-[265px] flex flex-row items-center justify-center">
          <NoResultsFound
            title={null}
            subtitle="No ports are exposed currently"
            compact
            image={<SmileySad size={32} weight={1} />}
            shadow={false}
            border={false}
          />
        </div>
      )}
    </div>
  );
};

export const ExposedPorts = ({
  ports,
  onChange,
}: {
  ports: IExposedPorts[];
  onChange: (ports: IExposedPorts[]) => void;
}) => {
  const { errors, handleChange, submit, values, resetValues } = useForm({
    initialValues: {
      port: '',
      targetPort: '',
    },
    validationSchema: Yup.object({
      port: Yup.number()
        .required()
        .test('is-valid', 'Port already exists.', (value) => {
          return !ports.some((p) => p.port === value);
        }),
      targetPort: Yup.number().min(0).max(65535).required(),
    }),
    onSubmit: (val) => {
      onChange?.([
        ...ports,
        {
          port:
            typeof val.port === 'string' ? parseInt(val.port, 10) : val.port,
          targetPort:
            typeof val.targetPort === 'string'
              ? parseInt(val.targetPort, 10)
              : val.targetPort,
        },
      ]);
      resetValues();
    },
  });

  return (
    <>
      <div className="flex flex-col gap-3xl">
        <div className="flex flex-row gap-3xl items-start">
          <div className="flex-1">
            <NumberInput
              label={
                <InfoLabel label="Expose Port" info="info about expose port" />
              }
              size="lg"
              error={!!errors.port}
              message={errors.port}
              value={values.port}
              onChange={({ target }) => {
                handleChange('port')(dummyEvent(parseValue(target.value, 0)));
              }}
            />
          </div>
          <div className="flex-1">
            <NumberInput
              min={0}
              max={65536}
              label={
                <InfoLabel info="info about target port" label="Target port" />
              }
              size="lg"
              autoComplete="off"
              value={values.targetPort}
              onChange={({ target }) => {
                handleChange('targetPort')(
                  dummyEvent(parseValue(target.value, 0))
                );
              }}
            />
          </div>
          <div className="flex pt-5xl">
            <IconButton
              icon={<Plus />}
              variant="basic"
              disabled={!values.port || !values.targetPort}
              onClick={submit}
            />
          </div>
        </div>
      </div>
      <ExposedPortList
        exposedPorts={ports}
        onDelete={(ep) => {
          onChange?.(ports.filter((v) => v.port !== ep.port));
        }}
      />
    </>
  );
};

export const QRCodeView = ({ data }: { data: string }) => {
  return (
    <div className="flex flex-row gap-7xl">
      <div className="flex flex-col gap-2xl">
        <div className="bodyLg-medium text-text-default">
          Use WireGuard on your phone
        </div>
        <ul className="flex flex-col gap-lg bodyMd text-text-default list-disc list-outside pl-2xl">
          <li>Download the app from Google Play or Apple Store</li>
          <li>Open the app on your Phone</li>
          <li>Tab on the ➕ Plus icon</li>
          <li>Point your phone to this screen to capture the QR code</li>
        </ul>
      </div>
      <div>
        <QRCode value={data} />
      </div>
    </div>
  );
};

export const decodeConfig = ({
  encoding,
  value,
}: {
  encoding: string;
  value: string;
}) => {
  switch (encoding) {
    case 'base64':
      return atob(value);
    default:
      return value;
  }
};

const downloadConfig = ({
  filename,
  data,
}: {
  filename: string;
  data: string;
}) => {
  downloadFile({ filename, data, format: 'text/plain' });
};

export const ShowWireguardConfig = ({
  visible,
  setVisible,
  deviceName,
  creationMethod,
}: // data,
{
  visible: boolean;
  setVisible: (visible: boolean) => void;
  deviceName: string;
  creationMethod: string;
}) => {
  const [mode, setMode] = useState<'config' | 'qr'>('config');

  const [data, setData] = useState<{
    value: string;
    encoding: string;
  }>();

  const api = useConsoleApi();

  const { data: devData, isLoading } = useCustomSwr(
    () => (deviceName ? `device-${deviceName}` : null),
    async () => api.getGlobalVpnDevice({ deviceName, gvpn: 'default' }),
    true
  );

  useEffect(() => {
    setData(devData?.wireguardConfig);
  }, [devData]);

  const modeView = () => {
    if (isLoading) {
      return (
        <div className="flex flex-col items-center justify-center">
          <LoadingPlaceHolder />
        </div>
      );
    }

    if (!data) {
      return (
        <div className="h-[100px] flex items-center justify-center">
          No wireguard config found.
        </div>
      );
    }

    const config = decodeConfig(data);
    switch (mode) {
      case 'qr':
        return <QRCodeView data={config} />;
      case 'config':
      default:
        return (
          <div className="flex flex-col gap-3xl">
            <div className="bodyMd text-text-default">
              Please use the following configuration to set up your WireGuard
              client.
            </div>
            <CodeView
              data={config}
              showShellPrompt={false}
              isMultilineData
              copy
            />
          </div>
        );
    }
  };

  return (
    <Popup.Root show={visible} onOpenChange={setVisible}>
      <Popup.Header>
        {creationMethod === 'kl'
          ? 'KL Managed Device'
          : mode === 'config'
          ? 'Wireguard Config'
          : 'Wireguard Config QR Code'}
      </Popup.Header>
      <Popup.Content>
        {creationMethod === 'kl' ? (
          <div className="flex flex-col gap-2xl">
            <div className="bodyLg-medium text-text-default">
              This device is managed by Kloudlite command line.
            </div>
            <div className="flex flex-col gap-lg bodyMd text-text-default list-disc list-outside pl-2xl" />
            <Button
              to="https://github.com/kloudlite/kl"
              content="You can install Kloudlite cli tool."
              variant="plain"
              linkComponent={Link}
            />
          </div>
        ) : (
          <div className="flex flex-col gap-xl">
            <ExtendedFilledTab
              value={mode}
              onChange={(v) => {
                setMode(v as any);
              }}
              items={[
                {
                  label: 'Config',
                  value: 'config',
                },
                {
                  label: 'QR Code',
                  value: 'qr',
                },
              ]}
            />
            {modeView()}
          </div>
        )}
      </Popup.Content>
      <Popup.Footer>
        {creationMethod === 'kl' ? (
          <Popup.Button closable content="Close" />
        ) : (
          <Popup.Button
            onClick={() => {
              if (!data) {
                toast.error('No wireguard config found.');
                return;
              }

              downloadConfig({
                filename: `wireguardconfig.conf`,
                data: decodeConfig(data),
              });
            }}
            content="Export"
            prefix={<ArrowLineDown />}
            variant="primary"
          />
        )}
      </Popup.Footer>
    </Popup.Root>
  );
};
