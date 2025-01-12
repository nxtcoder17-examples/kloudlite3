/* eslint-disable jsx-a11y/control-has-associated-label */
import { CopySimple } from '~/console/components/icons';
import { defer } from '@remix-run/node';
import { useLoaderData, useNavigate, useOutletContext } from '@remix-run/react';
import { ReactNode, useEffect, useState } from 'react';
import { Button } from '@kloudlite/design-system/atoms/button';
import { TextInput } from '@kloudlite/design-system/atoms/input';
import Select from '@kloudlite/design-system/atoms/select';
import { toast } from '@kloudlite/design-system/molecule/toast';
import {
  Box,
  DeleteContainer,
} from '~/console/components/common-console-components';
import { LoadingComp, pWrapper } from '~/console/components/loading-component';
import { awsRegions } from '~/console/dummy/consts';
import { useConsoleApi } from '~/console/server/gql/api-provider';
import { ICluster } from '~/console/server/gql/queries/cluster-queries';
import {
  ConsoleApiType,
  GQLServerHandler,
} from '~/console/server/gql/saved-queries';
import {
  ExtractNodeType,
  parseName,
  parseNodes,
} from '~/console/server/r-utils/common';
import { ensureAccountSet } from '~/console/server/utils/auth-utils';
import { getPagination } from '~/console/server/utils/common';
import useClipboard from '~/root/lib/client/hooks/use-clipboard';
import useForm from '~/root/lib/client/hooks/use-form';
import { useUnsavedChanges } from '~/root/lib/client/hooks/use-unsaved-changes';
import Yup from '~/root/lib/server/helpers/yup';
import { IRemixCtx } from '~/root/lib/types/common';
import { handleError } from '~/root/lib/utils/common';
import { mapper } from '@kloudlite/design-system/utils';
import DeleteDialog from '~/console/components/delete-dialog';
import { useReload } from '~/root/lib/client/helpers/reloader';
import { IProviderSecrets } from '~/console/server/gql/queries/provider-secret-queries';
import Wrapper from '~/console/components/wrapper';
import { IClusterContext } from '../../_layout';

export const loader = async (ctx: IRemixCtx) => {
  const promise = pWrapper(async () => {
    ensureAccountSet(ctx);
    const { data, errors } = await GQLServerHandler(
      ctx.request
    ).listProviderSecrets({
      pagination: getPagination(ctx),
    });

    if (errors) {
      throw errors[0];
    }

    return {
      providerSecrets: data,
    };
  });
  return defer({ promise });
};

export const updateCluster = async ({
  api,
  data,
}: {
  api: ConsoleApiType;
  data: ICluster;
}) => {
  try {
    const { errors: e } = await api.updateCluster({
      cluster: {
        displayName: data.displayName,
        metadata: {
          name: data.metadata.name,
        },
        spec: {
          cloudProvider: data.spec.cloudProvider,
          availabilityMode: data.spec.availabilityMode,
        },
      },
    });
    if (e) {
      throw e[0];
    }
  } catch (err) {
    handleError(err);
  }
};

const Layout = ({
  providerSecrets,
}: {
  providerSecrets: {
    label: string;
    value: string;
    render: () => ReactNode;
    provider: ExtractNodeType<IProviderSecrets>;
  }[];
}) => {
  const { account, cluster } = useOutletContext<IClusterContext>();
  const [deleteCluster, setDeleteCluster] = useState(false);
  const { setHasChanges, resetAndReload } = useUnsavedChanges();
  const reload = useReload();
  const api = useConsoleApi();
  const navigate = useNavigate();

  const { copy } = useClipboard({
    onSuccess() {
      toast.success('Text copied to clipboard.');
    },
  });

  const { values, handleChange, submit, isLoading, resetValues, setValues } =
    useForm({
      initialValues: {
        displayName: cluster.displayName,
      },
      validationSchema: Yup.object({
        displayName: Yup.string().required('Name is required.'),
      }),
      onSubmit: async (val) => {
        await updateCluster({
          api,
          data: { ...cluster, displayName: val.displayName },
        });
        resetAndReload();
      },
    });

  useEffect(() => {
    setValues({
      displayName: cluster.displayName,
    });
  }, [cluster]);

  useEffect(() => {
    setHasChanges(values.displayName !== cluster.displayName);
  }, [values]);

  // const defaultProvider = providerSecrets.find((ps) => cluster.spec?.aws);

  const defaultRegion = awsRegions.find(
    (r) => r.Name === cluster.spec?.aws?.region
  );
  return (
    <Wrapper
      secondaryHeader={{
        title: 'General',
        action: values.displayName !== cluster.displayName && (
          <div className="flex flex-row gap-3xl items-center">
            <Button
              content="Discard"
              variant="basic"
              onClick={() => {
                resetValues();
              }}
            />
            <Button
              content="Save changes"
              variant="primary"
              onClick={() => {
                submit();
              }}
              loading={isLoading}
            />
          </div>
        ),
      }}
    >
      <div className="flex flex-col gap-6xl">
        <Box title="General">
          <div className="flex flex-row items-center gap-3xl">
            <div className="flex-1">
              <TextInput
                label="Cluster name"
                value={values.displayName}
                onChange={handleChange('displayName')}
              />
            </div>
            <div className="flex-1">
              <TextInput
                value={cluster.metadata.name}
                label="Cluster ID"
                suffix={
                  <div
                    className="flex justify-center items-center"
                    title="Copy"
                  >
                    <button
                      onClick={() => copy(cluster.metadata.name)}
                      className="outline-none hover:bg-surface-basic-hovered active:bg-surface-basic-active rounded text-text-default"
                      tabIndex={-1}
                    >
                      <CopySimple size={16} />
                    </button>
                  </div>
                }
                disabled
              />
            </div>
          </div>
          <div className="flex flex-row items-center gap-3xl">
            {/* <div className="flex-1"> */}
            {/*   {' '} */}
            {/*   <Select */}
            {/*     label="Cloud Provider" */}
            {/*     placeholder="Select cloud provider" */}
            {/*     disabled */}
            {/*     value={defaultProvider?.value} */}
            {/*     options={async () => providerSecrets} */}
            {/*   /> */}
            {/* </div> */}
            <div className="flex-1">
              <Select
                disabled
                label="Region"
                placeholder="Select region"
                value={defaultRegion?.Name}
                options={async () =>
                  mapper(awsRegions, (v) => {
                    return {
                      value: v.Name,
                      label: v.Name,
                      region: v,
                    };
                  })
                }
              />
            </div>
          </div>
        </Box>

        <DeleteContainer
          title="Delete Cluster"
          action={async () => {
            setDeleteCluster(true);
          }}
        >
          Permanently remove your Cluster and all of its contents from the
          Kloudlite platform. This action is not reversible — please continue
          with caution.
        </DeleteContainer>
      </div>
      <DeleteDialog
        resourceName={parseName(cluster)}
        resourceType="cluster"
        show={deleteCluster}
        setShow={setDeleteCluster}
        onSubmit={async () => {
          try {
            const { errors } = await api.deleteCluster({
              name: parseName(cluster),
            });

            if (errors) {
              throw errors[0];
            }
            reload();
            toast.success(`Cluster deleted successfully`);
            setDeleteCluster(false);
            navigate(`/${parseName(account)}/infra/clusters`);
          } catch (err) {
            handleError(err);
          }
        }}
      />
    </Wrapper>
  );
};

const SettingGeneral = () => {
  const { promise } = useLoaderData<typeof loader>();
  return (
    <LoadingComp data={promise}>
      {({ providerSecrets }) => {
        const providerSecretsOptions = parseNodes(providerSecrets).map(
          (provider) => ({
            value: parseName(provider),
            label: provider.displayName,
            render: () => (
              <div className="flex flex-col">
                <div>{provider.displayName}</div>
                <div className="bodySm text-text-soft">
                  {parseName(provider)}
                </div>
              </div>
            ),
            provider,
          })
        );
        return <Layout providerSecrets={providerSecretsOptions} />;
      }}
    </LoadingComp>
  );
};
export default SettingGeneral;
