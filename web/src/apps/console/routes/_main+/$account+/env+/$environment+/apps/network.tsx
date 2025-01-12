import { useEffect } from 'react';
import { IconButton } from '@kloudlite/design-system/atoms/button';
import { usePagination } from '@kloudlite/design-system/molecule/pagination';
import { cn } from '@kloudlite/design-system/utils';
import List from '~/console/components/list';
import NoResultsFound from '~/console/components/no-results-found';
import {
  ChevronLeft,
  ChevronRight,
  SmileySad,
} from '~/console/components/icons';
import { IApp } from '~/console/server/gql/queries/app-queries';
import { NN } from '~/root/lib/types/common';
import { ExtractArrayType, parseValue } from '~/console/page-components/util';
import { ISetState } from '~/console/page-components/app-states';
import { NumberInput } from '@kloudlite/design-system/atoms/input';

export type exposedPortsType = ExtractArrayType<
  NN<NN<IApp['spec']['intercept']>['portMappings']>
>;

interface IExposedPortList {
  exposedPorts: exposedPortsType[];
  setExposedPorts: ISetState<exposedPortsType[]>;
}
const ExposedPortList = ({
  exposedPorts,
  setExposedPorts,
}: IExposedPortList) => {
  const { page, hasNext, hasPrevious, onNext, onPrev, setItems } =
    usePagination({
      items: exposedPorts,
      itemsPerPage: 5,
    });

  const updateDevPort = (p: exposedPortsType) => {
    setExposedPorts((s) => s.map((d) => (p.appPort === d.appPort ? p : d)));
  };

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
                key={ep.appPort}
                columns={[
                  {
                    key: `${index}-column-1`,
                    className: 'flex-1',
                    render: () => (
                      <div className="flex flex-row gap-md items-center bodyMd text-text-soft">
                        <span>Container: </span>
                        {ep.appPort}
                      </div>
                    ),
                  },
                  {
                    key: `${index}-column-2`,
                    className: 'flex-1',
                    render: () => (
                      <div className="flex flex-row gap-md items-center bodyMd text-text-soft">
                        <span>Device: </span>
                        <NumberInput
                          value={ep.devicePort}
                          onChange={(value) =>
                            updateDevPort({
                              devicePort: parseValue(value.target.value, 0),
                              appPort: ep.appPort,
                            })
                          }
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

export default ExposedPortList;
