import { defer } from '@remix-run/node';
import {
  useLoaderData,
  useNavigate,
  useOutletContext,
  useParams,
} from '@remix-run/react';
import { FormEventHandler, useCallback, useEffect, useState } from 'react';
import { NumberInput, TextInput } from '@kloudlite/design-system/atoms/input';
import Select from '@kloudlite/design-system/atoms/select';
import { Switch } from '@kloudlite/design-system/atoms/switch';
import { titleCase, useMapper } from '@kloudlite/design-system/utils';
import {
  BottomNavigation,
  ReviewComponent,
} from '~/console/components/commons';
import { LoadingComp, pWrapper } from '~/console/components/loading-component';
import MultiStepProgress, {
  useMultiStepProgress,
} from '~/console/components/multi-step-progress';
import MultiStepProgressWrapper from '~/console/components/multi-step-progress-wrapper';
import { NameIdView } from '~/console/components/name-id-view';
import { IAccountContext } from '~/console/routes/_main+/$account+/_layout';
import { IManagedServiceContext } from '~/console/routes/_main+/$account+/msvc+/$msv+/_layout';
import { useConsoleApi } from '~/console/server/gql/api-provider';
import { IClusterMSvs } from '~/console/server/gql/queries/cluster-managed-services-queries';
import { IMSvTemplate } from '~/console/server/gql/queries/managed-templates-queries';
import { GQLServerHandler } from '~/console/server/gql/saved-queries';
import {
  ExtractNodeType,
  parseName,
  parseNodes,
} from '~/console/server/r-utils/common';
import { getManagedTemplate } from '~/console/utils/commons';
import useForm, { dummyEvent } from '~/lib/client/hooks/use-form';
import Yup from '~/lib/server/helpers/yup';
import { IRemixCtx } from '~/lib/types/common';
import { handleError } from '~/lib/utils/common';

export const loader = (ctx: IRemixCtx) => {
  const promise = pWrapper(async () => {
    const { data: mData, errors: mErrors } = await GQLServerHandler(
      ctx.request
    ).listClusterMSvs({
      pagination: {
        orderBy: 'updateTime',
        sortDirection: 'DESC',
        first: 100,
      },
    });

    if (mErrors) {
      throw mErrors[0];
    }
    return { managedServicesData: mData };
  });
  return defer({ promise });
};

const RenderField = ({
  field,
  value,
  onChange,
  error,
  message,
}: {
  field: IMSvTemplate['fields'][number];
  onChange: (e: string) => (e: { target: { value: any } }) => void;
  value: any;
  error: boolean;
  message?: string;
}) => {
  const [qos, setQos] = useState(false);
  if (field.inputType === 'Number') {
    return (
      <NumberInput
        error={error}
        message={message}
        label={`${field.label}${field.required ? ' *' : ''}`}
        min={field.min}
        max={field.max}
        placeholder={field.label}
        value={parseFloat(value) / (field.multiplier || 1) || ''}
        onChange={({ target }) => {
          onChange(`res.${field.name}`)(
            dummyEvent(
              `${parseFloat(target.value) * (field.multiplier || 1)}${
                field.unit
              }`
            )
          );
        }}
        suffix={field.displayUnit}
      />
    );
  }

  if (field.inputType === 'String') {
    return (
      <TextInput
        label={field.label}
        value={value || ''}
        onChange={onChange(`res.${field.name}`)}
        suffix={field.displayUnit}
      />
    );
  }
  if (field.inputType === 'Resource') {
    return (
      <div className="flex flex-col gap-md">
        <div className="bodyMd-medium text-text-default">{`${field.label}${
          field.required ? ' *' : ''
        }`}</div>
        <div className="flex flex-row gap-xl items-center">
          <div className="flex flex-row gap-xl items-end flex-1 ">
            <div className="flex-1">
              <NumberInput
                error={error}
                message={message}
                min={field.min}
                max={field.max}
                placeholder={qos ? field.label : `${field.label} min`}
                value={parseFloat(value.min) / (field.multiplier || 1) || ''}
                onChange={({ target }) => {
                  onChange(`res.${field.name}.min`)(
                    dummyEvent(
                      `${parseFloat(target.value) * (field.multiplier || 1)}${
                        field.unit
                      }`
                    )
                  );
                  if (qos) {
                    onChange(`res.${field.name}.max`)(
                      dummyEvent(
                        `${parseFloat(target.value) * (field.multiplier || 1)}${
                          field.unit
                        }`
                      )
                    );
                  }
                }}
                suffix={field.displayUnit}
              />
            </div>
            {!qos && (
              <div className="flex-1">
                <NumberInput
                  error={error}
                  message={message}
                  min={field.min}
                  max={field.max}
                  placeholder={`${field.label} max`}
                  value={parseFloat(value.max) / (field.multiplier || 1)}
                  onChange={({ target }) => {
                    onChange(`res.${field.name}.max`)(
                      dummyEvent(
                        `${parseFloat(target.value) * (field.multiplier || 1)}${
                          field.unit
                        }`
                      )
                    );
                  }}
                  suffix={field.displayUnit}
                />
              </div>
            )}
          </div>
          <div className="flex flex-col gap-md min-w-[115px]">
            <Switch
              label="Guaranteed"
              checked={qos}
              onChange={(_value) => {
                setQos(_value);
                if (_value) {
                  onChange(`res.${field.name}.max`)(dummyEvent(`${value.min}`));
                }
              }}
            />
          </div>
        </div>
      </div>
    );
  }
  return <div>unknown input type {field.inputType}</div>;
};

const flatM = (obj: Record<string, any>) => {
  const flatJson = {};
  // eslint-disable-next-line guard-for-in
  for (const key in obj) {
    const parts = key.split('.');

    let temp: Record<string, any> = flatJson;

    if (parts.length === 1) {
      temp[key] = null;
    } else {
      parts.forEach((part, index) => {
        if (index === parts.length - 1) {
          temp[part] = {
            min: null,
            max: null,
          };
        } else {
          temp[part] = temp[part] || {};
        }
        temp = temp[part];
      });
    }
  }

  return flatJson;
};

type ISelectedResource = {
  label: string;
  value: string;
  resource: IMSvTemplate['resources'][number];
};

interface ITemplateView {
  handleSubmit: FormEventHandler<HTMLFormElement>;
  values: Record<string, any>;
  errors: Record<string, any>;
  resources: {
    label: string;
    value: string;
    resource: ExtractNodeType<IMSvTemplate>['resources'][number];
  }[];
  services: ExtractNodeType<IClusterMSvs>[];
  isLoading: boolean;
  handleChange: (key: string) => (e: { target: { value: any } }) => void;
}

const TemplateView = ({
  handleSubmit,
  values,
  handleChange,
  errors,
  services,
  resources,
  isLoading,
}: ITemplateView) => {
  return (
    <form className="flex flex-col gap-3xl" onSubmit={handleSubmit}>
      <div className="bodyMd text-text-soft">Create your managed resource.</div>
      <Select
        label="Resource type"
        size="lg"
        placeholder="Select resource type"
        value={values.selectedResource?.value}
        searchable
        onChange={(val) => {
          handleChange('selectedResource')(dummyEvent(val));
        }}
        options={async () => [...resources]}
        error={!!errors.selectedResource}
        message={errors.selectedResource}
      />
      <BottomNavigation
        primaryButton={{
          loading: isLoading,
          variant: 'primary',
          content: 'Next',
          type: 'submit',
        }}
      />
    </form>
  );
};

const FieldView = ({
  selectedResource,
  values,
  handleSubmit,
  handleChange,
  errors,
}: {
  handleChange: (key: string) => (e: { target: { value: any } }) => void;
  handleSubmit: FormEventHandler<HTMLFormElement>;
  values: Record<string, any>;
  errors: Record<string, any>;
  selectedResource: ISelectedResource | null;
}) => {
  return (
    <form
      className="flex flex-col gap-3xl"
      onSubmit={(e) => {
        if (!values.isNameError) {
          handleSubmit(e);
        } else {
          e.preventDefault();
        }
      }}
    >
      <NameIdView
        placeholder="Enter managed resource name"
        label="Name"
        resType="managed_resource"
        name={values.name}
        displayName={values.displayName}
        errors={errors.name}
        handleChange={handleChange}
        nameErrorLabel="isNameError"
      />
      {selectedResource?.resource?.fields?.map((field) => {
        const k = field.name;
        const x = k.split('.').reduce((acc, curr) => {
          if (!acc) {
            return values.res?.[curr];
          }
          return acc[curr];
        }, null);

        return (
          <RenderField
            field={field}
            key={field.name}
            onChange={handleChange}
            value={x}
            error={!!errors[k]}
            message={errors[k]}
          />
        );
      })}
      <BottomNavigation
        primaryButton={{
          variant: 'primary',
          content: 'Next',
          type: 'submit',
        }}
      />
    </form>
  );
};

const ReviewView = ({
  handleSubmit,
  values,
  selectedResource,
  selectedService: managedService,
  isLoading,
  onEdit,
}: {
  values: Record<string, any>;
  handleSubmit: FormEventHandler<HTMLFormElement>;
  selectedResource: ISelectedResource | null;
  selectedService: IManagedServiceContext['managedService'] | null;
  isLoading?: boolean;
  onEdit: (step: number) => void;
}) => {
  const renderFieldView = () => {
    const fields = Object.entries(values.res).filter(
      ([k, _v]) => !['resources'].includes(k)
    );
    if (fields.length > 0) {
      return (
        <ReviewComponent
          title="Fields"
          onEdit={() => {
            onEdit(2);
          }}
        >
          <div className="flex flex-col p-xl  gap-lg rounded border border-border-default flex-1 overflow-hidden">
            {fields?.map(([key, value]) => {
              const k = key as string;
              const v = value as string;
              return (
                <div
                  key={k}
                  className="flex flex-col gap-md  [&:not(:last-child)]:pb-lg   [&:not(:last-child)]:border-b border-border-default"
                >
                  <div className="bodyMd-medium text-text-default">
                    {titleCase(k)}
                  </div>
                  <div className="bodySm text-text-soft">{v}</div>
                </div>
              );
            })}
          </div>
        </ReviewComponent>
      );
    }
    return null;
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-3xl">
      <div className="flex flex-col gap-xl">
        <ReviewComponent
          title="Basic detail"
          onEdit={() => {
            onEdit(2);
          }}
        >
          <div className="flex flex-col p-xl gap-lg rounded border border-border-default">
            <div className="flex flex-col gap-md">
              <div className="bodyMd-semibold text-text-default">
                {values.displayName}
              </div>
              <div className="bodySm text-text-soft">{values.name}</div>
            </div>
          </div>
        </ReviewComponent>
        {selectedResource && (
          <ReviewComponent
            title="Service detail"
            onEdit={() => {
              onEdit(1);
            }}
          >
            <div className="flex flex-col p-xl gap-lg rounded border border-border-default">
              <div className="flex flex-col gap-md pb-lg border-b border-border-default">
                <div className="bodyMd-semibold text-text-default">Service</div>
                <div className="bodySm text-text-soft">
                  {managedService?.metadata?.name}
                </div>
              </div>
              <div className="flex flex-col gap-md">
                <div className="bodyMd-semibold text-text-default">
                  Resource type
                </div>
                <div className="bodySm text-text-soft">
                  {selectedResource?.resource?.name}
                </div>
              </div>
            </div>
          </ReviewComponent>
        )}
        {renderFieldView()}
      </div>
      <BottomNavigation
        primaryButton={{
          variant: 'primary',
          content: 'Create',
          loading: isLoading,
          type: 'submit',
        }}
      />
    </form>
  );
};

const App = ({ services }: { services: ExtractNodeType<IClusterMSvs>[] }) => {
  const { msvtemplates } = useOutletContext<IAccountContext>();
  const navigate = useNavigate();
  const api = useConsoleApi();

  const { account, msv } = useParams();
  const rootUrl = `/${account}/msvc/${msv}/managed-resources`;
  const { currentStep, jumpStep, nextStep } = useMultiStepProgress({
    defaultStep: 1,
    totalSteps: 3,
  });

  const { managedService } = useOutletContext<IManagedServiceContext>();

  const commonTemplates = useCallback(() => {
    return getManagedTemplate({
      templates: msvtemplates || [],
      kind: managedService?.spec?.msvcSpec.serviceTemplate.kind || '',
      apiVersion:
        managedService?.spec?.msvcSpec.serviceTemplate.apiVersion || '',
    });
  }, [managedService, msvtemplates]);

  const { values, errors, handleSubmit, handleChange, isLoading, setValues } =
    useForm({
      initialValues: {
        name: '',
        displayName: '',
        selectedResource: (() => {
          const ct = commonTemplates()?.resources;
          if (ct && ct.length === 1) {
            return {
              label: ct[0].displayName || '',
              value: ct[0].name || '',
              resource: ct[0],
            };
          }
          return null;
        })(),
        res: {},
        isNameError: false,
      },
      validationSchema: Yup.object({
        name: Yup.string().test('required', 'Name is required', (v) => {
          return !(currentStep === 2 && !v);
        }),
        displayName: Yup.string().test('required', 'Name is required', (v) => {
          return !(currentStep === 2 && !v);
        }),
        selectedResource: Yup.object({}).required('Resource type is required'),
      }),
      onSubmit: async (val) => {
        const selectedResource =
          val.selectedResource as unknown as ISelectedResource;

        const submit = async () => {
          try {
            if (!msv) {
              throw new Error('msvc is required!.');
            }
            if (
              !managedService ||
              (managedService && !managedService.spec?.msvcSpec.serviceTemplate)
            ) {
              throw new Error('Service apiversion or kind error.');
            }
            const { errors: e } = await api.createManagedResource({
              msvcName: msv,
              mres: {
                displayName: val.displayName,
                metadata: {
                  name: val.name,
                },

                spec: {
                  resourceTemplate: {
                    apiVersion: selectedResource.resource.apiVersion || '',
                    kind: selectedResource.resource.kind || '',
                    spec: {
                      ...val.res,
                    },
                    msvcRef: {
                      name: parseName(managedService),
                      namespace: managedService?.spec?.targetNamespace || '',
                      apiVersion:
                        managedService?.spec?.msvcSpec.serviceTemplate
                          .apiVersion || '',
                      kind:
                        managedService?.spec?.msvcSpec.serviceTemplate.kind ||
                        '',
                      // clusterName: managedService.clusterName,
                    },
                  },
                },
              },
            });
            if (e) {
              throw e[0];
            }
            navigate(rootUrl);
          } catch (err) {
            handleError(err);
          }
        };
        switch (currentStep) {
          case 1:
            nextStep();
            break;
          case 2:
            nextStep();
            break;
          case 3:
            await submit();
            break;
          default:
            break;
        }
      },
    });

  useEffect(() => {
    const selectedResource =
      values?.selectedResource as unknown as ISelectedResource;
    if (selectedResource?.resource?.fields) {
      setValues({
        ...values,
        res: {
          ...flatM(
            selectedResource?.resource?.fields.reduce((acc, curr) => {
              return { ...acc, [curr.name]: curr.defaultValue };
            }, {})
          ),
        },
      });
    }
  }, [values.selectedResource]);

  const resources = useMapper(
    [...(commonTemplates()?.resources || [])],
    (res) => ({
      label: res.displayName,
      value: res.name,
      resource: res,
    })
  );

  return (
    <MultiStepProgressWrapper
      title="Let’s create new managed resource."
      subTitle="Simplify Collaboration and Enhance Productivity with Kloudlite teams"
      backButton={{
        content: 'Back to managed resources',
        to: rootUrl,
      }}
    >
      <MultiStepProgress.Root currentStep={currentStep} jumpStep={jumpStep}>
        <MultiStepProgress.Step label="Select service" step={1}>
          <TemplateView
            isLoading={isLoading}
            services={services}
            handleChange={handleChange}
            handleSubmit={handleSubmit}
            errors={errors}
            values={values}
            resources={resources}
          />
        </MultiStepProgress.Step>
        <MultiStepProgress.Step label="Configure managed resource" step={2}>
          <FieldView
            selectedResource={values.selectedResource}
            values={values}
            errors={errors}
            handleChange={handleChange}
            handleSubmit={handleSubmit}
          />
        </MultiStepProgress.Step>
        <MultiStepProgress.Step label="Review" step={3}>
          <ReviewView
            onEdit={jumpStep}
            values={values}
            handleSubmit={handleSubmit}
            selectedService={managedService}
            selectedResource={values.selectedResource}
            isLoading={isLoading}
          />
        </MultiStepProgress.Step>
      </MultiStepProgress.Root>
    </MultiStepProgressWrapper>
  );
};

const ManagedServiceLayout = () => {
  const { promise } = useLoaderData<typeof loader>();
  return (
    <LoadingComp data={promise}>
      {({ managedServicesData }) => {
        const managedServicesList = parseNodes(managedServicesData);
        return <App services={managedServicesList} />;
      }}
    </LoadingComp>
  );
};

const NewManagedService = () => {
  return <ManagedServiceLayout />;
};

export const handle = {
  noMainLayout: true,
};

export default NewManagedService;
