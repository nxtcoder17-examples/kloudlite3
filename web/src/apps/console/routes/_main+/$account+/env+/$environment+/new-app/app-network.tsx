import { useEffect, useState } from 'react';
import { IconButton } from '@kloudlite/design-system/atoms/button';
// import { NumberInput } from '@kloudlite/design-system/atoms/input';
import { usePagination } from '@kloudlite/design-system/molecule/pagination';
import { cn } from '@kloudlite/design-system/utils';
import { BottomNavigation } from '~/console/components/commons';
import {
  ChevronLeft,
  ChevronRight,
  SmileySad,
  X,
} from '~/console/components/icons';
import List from '~/console/components/list';
import NoResultsFound from '~/console/components/no-results-found';
import { useAppState } from '~/console/page-components/app-states';
import { FadeIn } from '~/console/page-components/util';
import { useUnsavedChanges } from '~/lib/client/hooks/use-unsaved-changes';
// import { dummyEvent } from '~/root/lib/client/hooks/use-form';
import Select from '@kloudlite/design-system/atoms/select';

interface IExposedPorts {
  port: number;
}

interface IExposedPortList {
  exposedPorts: IExposedPorts[];
  onDelete: (exposedPorts: IExposedPorts) => void;
}
const ExposedPortList = ({
  exposedPorts,
  onDelete = (_) => _,
}: IExposedPortList) => {
  const { page, hasNext, hasPrevious, onNext, onPrev, setItems } =
    usePagination({
      items: exposedPorts,
      itemsPerPage: 5,
    });

  useEffect(() => {
    setItems(exposedPorts);
  }, [exposedPorts]);
  return (
    <div className="flex flex-col gap-lg bg-surface-basic-default">
      {exposedPorts.length > 0 && (
        <List.Root
          className="min-h-[347px] !shadow-none"
          header={
            <div className="flex flex-row items-center justify-between w-full">
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
                  '!border-b': index < 4,
                  '!rounded-b-none': index < 4,
                })}
                key={ep.port}
                columns={[
                  {
                    key: `${index}-column-2`,
                    className: 'flex-1',
                    render: () => (
                      <div className="flex flex-row gap-md items-center bodyMd text-text-soft">
                        <span>Container: </span>
                        {ep.port}
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
        <div className="rounded border-border-default border min-h-[347px] flex flex-row items-center justify-center">
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

const ExposedRoute = () => {
  const { app, setApp } = useAppState();

  return (
    <div className="flex-1">
      <Select
        creatable
        size="lg"
        label="Exposed Domains"
        multiple
        value={app.spec.router?.domains.map((s) => `${s}`)}
        options={async () =>
          (app.spec.router?.domains || []).map((s) => ({
            label: `${s}`,
            value: `${s}`,
          }))
        }
        onChange={(val, v) => {
          // const domains = app.spec.router?.domains || [];
          setApp({
            ...app,
            spec: {
              ...app.spec,
              router: {
                ...app.spec.router,
                domains: [...v],
              },
            },
          });
        }}
        // error={!!portError}
        // message={portError}
        disableWhileLoading
        createLabel="Exposed Domains"
      />
    </div>
  );
};

export const ExposedPorts = () => {
  const [portError, setPortError] = useState<string>('');

  const { services, setServices } = useAppState();

  // for updating
  const { hasChanges } = useUnsavedChanges();

  // for updating
  useEffect(() => {
    if (!hasChanges) {
      // setPort('');
      setPortError('');
    }
  }, [hasChanges]);

  return (
    <div className="flex flex-col gap-3xl ">
      <div className="flex flex-row gap-3xl items-start">
        <div className="flex-1">
          {/* <NumberInput
              min={0}
              max={65534}
              label={
                <InfoLabel label="Expose Port" info="info about expose port" />
              }
              size="lg"
              error={!!portError}
              message={portError}
              value={port}
              onChange={({ target }) => {
                setPort(parseValue(target.value, 0));
              }}
            /> */}
          <Select
            creatable
            size="lg"
            label="Exposed ports"
            multiple
            value={services.map((s) => `${s.port}`)}
            options={async () => []}
            onChange={(val, v) => {
              const r = /^\d+$/;
              if (v.every((c) => r.test(c))) {
                setServices([...v.map((vv) => ({ port: parseInt(vv, 10) }))]);
              } else {
                setServices((prev) => [...prev]);
              }
            }}
            error={!!portError}
            message={portError}
            disableWhileLoading
            createLabel="Exposed ports"
          />
        </div>
      </div>
      <div className="flex flex-row gap-md items-center">
        <div className="bodySm text-text-soft flex-1">
          All network entries be mounted on the path specified in the container
        </div>
        {/* <Button
            content="Expose port"
            variant="basic"
            disabled={!ports}
            onClick={() => {
              if (services?.find((ep) => ep.port && ep.port === port)) {
                setPortError('Port is already exposed.');
              } else {
                if (typeof port === 'number')
                  setServices((prev) => [
                    ...prev,
                    {
                      port,
                    },
                  ]);
                setPort('');
                setPortError('');
              }
            }}
          /> */}
      </div>
    </div>
  );
};

export const Network = () => {
  return (
    <div className="flex flex-col gap-3xl p-3xl rounded border border-border-default">
      <ExposedPorts />
      {/* <ExposedRoute /> */}
    </div>
  );
};

const AppNetwork = () => {
  const { setPage, markPageAsCompleted } = useAppState();
  return (
    <FadeIn notForm>
      <div className="bodyMd text-text-soft">
        Expose service ports that need to be exposed from container
      </div>
      <div className="flex flex-col gap-3xl p-3xl rounded border border-border-default">
        <ExposedPorts />
      </div>

      <BottomNavigation
        primaryButton={{
          type: 'submit',
          content: 'Save & Continue',
          variant: 'primary',
          onClick: () => {
            setPage(5);
            markPageAsCompleted(4);
            markPageAsCompleted(5);
          },
        }}
        secondaryButton={{
          content: 'Environments',
          variant: 'outline',
          onClick: () => {
            setPage(3);
          },
        }}
      />
    </FadeIn>
  );
};

export default AppNetwork;
