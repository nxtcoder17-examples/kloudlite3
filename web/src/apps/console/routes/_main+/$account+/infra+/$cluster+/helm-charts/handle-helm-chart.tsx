/* eslint-disable react/destructuring-assignment */
import { TextArea, TextInput } from '@kloudlite/design-system/atoms/input';
import Popup from '@kloudlite/design-system/molecule/popup';
import { IDialogBase } from '~/console/components/types.d';
import { useConsoleApi } from '~/console/server/gql/api-provider';
import { IHelmCharts } from '~/console/server/gql/queries/helm-chart-queries';
import {
  ExtractNodeType,
  parseName,
  parseNodes,
} from '~/console/server/r-utils/common';
import { useReload } from '~/root/lib/client/helpers/reloader';
import useForm, { dummyEvent } from '~/root/lib/client/hooks/use-form';
import Yup from '~/root/lib/server/helpers/yup';
import { handleError } from '~/root/lib/utils/common';
import yaml from 'js-yaml';
import { useParams } from '@remix-run/react';
import axios from 'axios';
import useDebounce from '~/root/lib/client/hooks/use-debounce';
import { ReactNode, useEffect, useRef, useState } from 'react';
import Select from '@kloudlite/design-system/atoms/select';
import { CircleWavyCheckFill } from '~/console/components/icons';
import { cn, useMapper } from '@kloudlite/design-system/utils';
import Pulsable from 'react-pulsable';
import { NameIdView } from '~/console/components/name-id-view';
import useCustomSwr from '~/root/lib/client/hooks/use-custom-swr';
import { toast } from '@kloudlite/design-system/molecule/toast';
import ExtendedFilledTab from '~/console/components/extended-filled-tab';
import { keyconstants } from '~/console/server/r-utils/key-constants';
import logger from '~/root/lib/client/helpers/log';

const LOGO_URL = 'https://artifacthub.io/image/';

type IDialog = IDialogBase<ExtractNodeType<IHelmCharts>>;

type IHelmDoc = {
  apiVersion: string;
  entries: {
    [key: string]: { version: string }[];
  };
  generated: string;
};

const repoRenderer = ({
  value,
  repoUrl,
}: {
  value: string;
  repoUrl: string;
}) => {
  return (
    <div className="flex flex-row gap-xl items-center bodyMd text-text-default">
      <span>{!repoUrl ? value : repoUrl}</span>
    </div>
  );
};

const filterUniqueVersions = (versions: IHelmDoc['entries']['keys']) => {
  return versions.filter(
    (obj, index, self) =>
      index === self.findIndex((t) => t.version === obj.version)
  );
};
const Root = (props: IDialog) => {
  const { isUpdate, setVisible } = props;
  const api = useConsoleApi();
  const reloadPage = useReload();
  const { cluster } = useParams();

  const [hemlCharts, setHelmCharts] = useState<
    Array<{ label: string; value: string; item: IHelmDoc['entries']['key'] }>
  >([]);

  const [chartVersions, setChartVersions] = useState<
    IHelmDoc['entries']['key']
  >([]);

  const [helmChartsLoading, setHelmChartsLoading] = useState(false);
  const [repoSearchText, setRepoSearchText] = useState('');
  const [repos, setRepos] = useState<
    {
      label: string;
      value: string;
      repoUrl: string;
      render: () => ReactNode;
    }[]
  >([]);

  const [reposLoading, setReposLoading] = useState(false);
  const [chartName, setChartName] = useState<
    { label: string; value: string } | undefined
  >(undefined);
  const [chartVersion, setChartVersion] = useState<
    { label: string; value: string } | undefined
  >(undefined);
  const [isRepoCreatable, setIsRepoCreatable] = useState(false);
  const [selectedRepo, setSelectedRepo] = useState<string>('');
  const [repoErrors, setRepoErrors] = useState(false);
  const [helmValues, setHelmValues] = useState('');
  const [activeTab, setActiveTab] = useState('defaults');

  const fetchValues = async ({
    packageId,
    version,
  }: {
    packageId: string;
    version: string;
  }) => {
    try {
      const r = await axios({
        method: 'get',
        url: `/artifacthub-values-api`,
        params: {
          packageId,
          version,
        },
      });
      setHelmValues(r.data);
    } catch (err) {
      toast.error('Error fetching chart values');
    }
  };

  useEffect(() => {
    if (isUpdate) {
      if (hemlCharts && hemlCharts.length > 0) {
        setChartVersions(
          hemlCharts.find((v) => v.value === props.data.spec?.chartName)
            ?.item || []
        );
        setChartName({
          label: props.data.spec?.chartName || '',
          value: props.data.spec?.chartName || '',
        });
      }
    } else {
      setChartName(undefined);
      setChartVersions([]);
    }
  }, [hemlCharts]);

  useEffect(() => {
    if (isUpdate) {
      if (chartVersions && chartVersions.length > 0) {
        setChartVersion({
          label: props.data.spec?.chartVersion || '',
          value: props.data.spec?.chartVersion || '',
        });
        if (props.data.spec?.chartVersion) {
          fetchValues({
            packageId:
              props.data.metadata?.annotations?.[
                keyconstants.helmChartRepoPackageId
              ],
            version: props.data.spec?.chartVersion,
          });
        }
      }
    } else {
      setChartVersion(undefined);
    }
  }, [chartVersions]);

  const {
    data: namespacesData,
    isLoading: namespacesIsLoading,
    error: namespacesError,
  } = useCustomSwr('/infra_namespaces', async () => {
    if (!cluster) {
      throw new Error('Cluster is required!.');
    }
    return api.listNamespaces({ clusterName: cluster });
  });

  const namespaces = useMapper(parseNodes(namespacesData), (val) => {
    return { label: parseName(val), value: parseName(val) };
  });

  const fetchHelmCharts = async (repoUrl: string) => {
    try {
      setRepoErrors(false);
      setHelmChartsLoading(true);
      const res = await axios.get(`/helmchart-api?url=${repoUrl}`);
      const repos = yaml.load(res.data, { json: true }) as IHelmDoc;
      setHelmCharts(
        Object.entries(repos.entries).map(([key, value]) => ({
          label: key,
          value: key,
          item: value,
        }))
      );
    } catch (error) {
      logger.log(error);
      setRepoErrors(true);
    } finally {
      setHelmChartsLoading(false);
    }
  };

  const searchRepos = async (text: string) => {
    setChartVersions([]);
    if (text) {
      try {
        const r = await axios({
          method: 'get',
          url: '/artifacthub-api',
          params: {
            offset: 0,
            limit: 10,
            kind: 0,
            ts_query_web: text,
          },
        });

        setRepos(
          r.data.packages.map(
            (hc: {
              name: string;
              package_id: string;
              logo_image_id: string;
              repository: {
                url: string;
                name: string;
                verified_publisher: boolean;
                organization_display_name?: string;
                user_alias?: string;
              };
            }) => ({
              label: hc.name,
              value: hc.package_id,
              repoUrl: hc.repository.url,
              render: () => (
                <div className="flex flex-row gap-xl items-center">
                  <Pulsable isLoading={!hc.logo_image_id} noPadding>
                    <span className=" pulsable pulsable-img">
                      <img
                        className={cn({
                          'w-4xl aspect-square object-contain': true,
                        })}
                        src={`${LOGO_URL}${hc.logo_image_id}`}
                        alt={hc.name}
                      />
                    </span>
                  </Pulsable>
                  <div className="flex flex-col flex-1">
                    <div className="flex flex-row gap-lg items-center">
                      <div className="flex-1">{hc.name}</div>
                      <div className="text-icon-primary mt-sm">
                        {hc.repository.verified_publisher && (
                          <CircleWavyCheckFill size={12} />
                        )}
                      </div>
                    </div>
                    <div className="bodySm text-text-disabled flex flex-row gap-md lowercase">
                      <span>
                        {hc.repository.organization_display_name ? (
                          <span>
                            ORG:{' '}
                            <span className="bodySm-semibold">
                              {hc.repository.organization_display_name}
                            </span>
                          </span>
                        ) : (
                          <span>
                            USER:{' '}
                            <span className="bodySm-semibold">
                              {hc.repository.user_alias}
                            </span>
                          </span>
                        )}
                      </span>{' '}
                      |{' '}
                      <span>
                        REPO:{' '}
                        <span className="bodySm-semibold">
                          {hc.repository.name}
                        </span>
                      </span>
                    </div>
                  </div>
                </div>
              ),
            })
          )
        );
      } catch {
        setRepoErrors(true);
      } finally {
        setReposLoading(false);
      }
    } else {
      //
      setReposLoading(false);
      setRepos([]);
    }
  };

  useDebounce(
    async () => {
      if (!repoSearchText.startsWith('https://')) {
        searchRepos(repoSearchText);
        setIsRepoCreatable(false);
      } else {
        setIsRepoCreatable(true);
        setReposLoading(false);
        setRepos([]);
      }
    },
    200,
    [repoSearchText]
  );

  const { values, errors, handleSubmit, handleChange, isLoading, resetValues } =
    useForm({
      initialValues: !isUpdate
        ? {
            displayName: '',
            name: '',
            namespace: '',
            chartName: '',
            chartRepoURL: '',
            values: '',
            isNameError: false,
          }
        : {
            isNameError: false,
            displayName: props.data.displayName,
            name: props.data.metadata?.name || '',
            values:
              Object.keys(props.data.spec?.values).length > 0
                ? yaml.dump(props.data.spec?.values)
                : '',
            namespace: props.data.metadata?.namespace,
            chartName: props.data.spec?.chartName,
            chartRepoURL: props.data.spec?.chartRepoURL,
          },
      validationSchema: Yup.object({
        displayName: Yup.string().required(),
        name: Yup.string().required(),
        namespace: Yup.string().required(),
        chartRepoURL: Yup.string(),
      }),

      onSubmit: async (val) => {
        if (!val.name) {
          throw new Error('This helm chart has no name');
        }
        if (!chartName?.value || !chartVersion?.value) {
          return;
        }
        if (!cluster) {
          throw new Error('Cluster is required.');
        }
        try {
          if (isUpdate) {
            if (!props.data || !props.data.spec) {
              throw new Error('No spec found');
            }
            const { errors } = await api.updateHelmChart({
              clusterName: cluster || '',
              release: {
                displayName: val.displayName,
                metadata: {
                  name: val.name,
                  namespace: val.namespace,
                  annotations: {
                    [keyconstants.helmChartRepoPackageId]:
                      props.data.metadata?.annotations?.[
                        keyconstants.helmChartRepoPackageId
                      ],
                  },
                },
                spec: {
                  chartName: props.data.spec?.chartName,
                  chartVersion: chartVersion?.value,
                  chartRepoURL: val.chartRepoURL!,
                  values: val.values
                    ? yaml.load(val.values, { json: true })
                    : {},
                },
              },
            });

            if (errors) {
              throw errors[0];
            }
          } else {
            const { errors } = await api.createHelmChart({
              clusterName: cluster || '',
              release: {
                displayName: val.displayName,
                metadata: {
                  name: val.name,
                  namespace: val.namespace,
                  annotations: {
                    [keyconstants.helmChartRepoPackageId]: selectedRepo,
                  },
                },
                spec: {
                  chartName: chartName?.value,
                  chartVersion: chartVersion?.value,
                  chartRepoURL: val.chartRepoURL!,
                  values: val.values
                    ? yaml.load(val.values, { json: true })
                    : {},
                },
              },
            });
            if (errors) {
              throw errors[0];
            }
          }

          reloadPage();
          setVisible(false);
          resetValues();
        } catch (error) {
          handleError(error);
        }
      },
    });

  useDebounce(
    () => {
      if (values.chartRepoURL) {
        fetchHelmCharts(values.chartRepoURL);
      }
    },
    300,
    [values.chartRepoURL]
  );

  const nameInputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    nameInputRef.current?.focus();
  }, []);

  return (
    <Popup.Form
      onSubmit={(e) => {
        if (!values.isNameError) {
          handleSubmit(e);
        } else {
          e.preventDefault();
        }
      }}
    >
      <Popup.Content className="!w-[900px]">
        <div className="flex flex-row gap-2xl ">
          <div className="flex flex-col gap-2xl basis-full border-border-default border-r pr-2xl">
            <NameIdView
              resType="helm_release"
              displayName={values.displayName}
              name={values.name}
              label="Name"
              placeholder="Enter helm chart name"
              errors={errors.name}
              handleChange={handleChange}
              nameErrorLabel="isNameError"
              isUpdate={isUpdate}
            />

            <Select
              label="Namespace"
              placeholder="Namespace"
              error={!!errors.namespace || (!isUpdate && !!namespacesError)}
              message={
                errors.namespace ||
                (!isUpdate && namespacesError
                  ? 'Error fetching namespaces'
                  : '')
              }
              disabled={isUpdate || namespacesIsLoading}
              options={async () =>
                isUpdate
                  ? [
                      {
                        label: values.namespace || '',
                        value: values.namespace || '',
                      },
                    ]
                  : namespaces
              }
              value={values.namespace}
              creatable={!isUpdate}
              onChange={(_, val) => {
                handleChange('namespace')(dummyEvent(val.toLowerCase()));
              }}
              noOptionMessage={
                <div className="p-2xl bodyMd text-center">
                  Type to create new namespace
                </div>
              }
            />

            {!isUpdate ? (
              <Select
                // open
                showclear={!!values.chartRepoURL}
                label="Chart repo url"
                placeholder="Search for or enter the repo url"
                searchable
                creatable={isRepoCreatable}
                options={async () => repos}
                value={selectedRepo}
                onChange={(value) => {
                  if (!repoSearchText.startsWith('https://')) {
                    handleChange('chartRepoURL')(dummyEvent(value.repoUrl));
                  } else {
                    handleChange('chartRepoURL')(dummyEvent(value.value));
                  }
                  setSelectedRepo(value.value);
                  setHelmCharts([]);
                }}
                onSearch={(text) => {
                  setRepoSearchText(text);
                  setReposLoading(true);
                }}
                valueRender={repoRenderer}
                loading={reposLoading}
                noOptionMessage={
                  <div className="p-2xl bodyMd text-center">
                    Search for or enter the repo url
                  </div>
                }
                error={repoErrors}
                message={repoErrors ? 'Error loading helm charts.' : ''}
              />
            ) : (
              <TextInput
                value={values.chartRepoURL}
                disabled
                label="Chart repo url"
              />
            )}
            {!isUpdate ? (
              <Select
                label="Chart name"
                placeholder="Chart name"
                searchable
                disabled={hemlCharts.length === 0 || reposLoading || repoErrors}
                value={chartName?.value}
                options={async () => hemlCharts}
                loading={!repoErrors && helmChartsLoading}
                onChange={(val) => {
                  setChartName(val);
                  setChartVersions(filterUniqueVersions(val.item));
                }}
                onSearch={() => true}
              />
            ) : (
              <TextInput
                value={chartName?.value}
                disabled
                label="Chart name"
                placeholder="Chart name"
              />
            )}
            <Select
              searchable
              label="Chart version"
              placeholder="Chart version"
              disabled={
                (!isUpdate &&
                  (chartVersions.length === 0 ||
                    reposLoading ||
                    helmChartsLoading ||
                    repoErrors)) ||
                (isUpdate && helmChartsLoading)
              }
              value={chartVersion?.value}
              options={async () => [
                ...chartVersions.map((cv) => ({
                  label: cv.version,
                  value: cv.version,
                })),
              ]}
              loading={isUpdate && helmChartsLoading}
              onChange={(val) => {
                setChartVersion(val);
                setHelmValues('');
                fetchValues({ packageId: selectedRepo, version: val.value });
              }}
              onSearch={() => true}
            />
          </div>
          <div className="basis-full flex flex-col">
            {chartVersion ? (
              <div className="flex flex-col gap-3xl h-full">
                <ExtendedFilledTab
                  value={activeTab}
                  onChange={setActiveTab}
                  items={[
                    { label: 'Defaults', value: 'defaults' },
                    {
                      label: 'Values',
                      value: 'values',
                    },
                  ]}
                />
                <TextArea
                  containerClassName="h-full"
                  className="h-full"
                  textFieldClassName={cn(
                    '!font-mono whitespace-pre break-normal overflow-x-scroll'
                  )}
                  placeholder={
                    activeTab === 'defaults' ? 'Default values' : 'Helm Values'
                  }
                  onChange={(e) => {
                    if (activeTab === 'values') {
                      handleChange('values')(e);
                    }
                  }}
                  error={!!errors.values}
                  message={errors.values}
                  value={(() => {
                    if (activeTab === 'defaults') {
                      return helmValues;
                    }
                    if (activeTab === 'values') {
                      return values.values;
                    }
                    return '';
                  })()}
                  name="helm-chart-values"
                />
              </div>
            ) : (
              <div className="flex items-center justify-center bodyMd flex-col h-full">
                Select chart name and version
              </div>
            )}
          </div>
        </div>
      </Popup.Content>
      <Popup.Footer>
        <Popup.Button content="Cancel" variant="basic" closable />
        <Popup.Button
          loading={isLoading}
          type="submit"
          content={isUpdate ? 'Update' : 'Install'}
          variant="primary"
        />
      </Popup.Footer>
    </Popup.Form>
  );
};

const HandleHelmChart = (props: IDialog) => {
  const { isUpdate, setVisible, visible } = props;

  return (
    <Popup.Root
      show={visible}
      onOpenChange={(v) => setVisible(v)}
      className="!w-[900px]"
    >
      <Popup.Header>
        {isUpdate ? 'Edit helm chart' : 'Install helm chart'}
      </Popup.Header>
      {(!isUpdate || (isUpdate && props.data)) && <Root {...props} />}
    </Popup.Root>
  );
};

export default HandleHelmChart;
