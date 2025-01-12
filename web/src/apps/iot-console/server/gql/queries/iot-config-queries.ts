import gql from 'graphql-tag';
import { IExecutor } from '~/root/lib/server/helpers/execute-query-with-context';
import { NN } from '~/root/lib/types/common';
import {
  ConsoleCreateConfigMutation,
  ConsoleCreateConfigMutationVariables,
  ConsoleGetConfigQuery,
  ConsoleGetConfigQueryVariables,
  ConsoleListConfigsQuery,
  ConsoleListConfigsQueryVariables,
  ConsoleUpdateConfigMutation,
  ConsoleUpdateConfigMutationVariables,
  ConsoleDeleteConfigMutation,
  ConsoleDeleteConfigMutationVariables,
} from '~/root/src/generated/gql/server';

export type IConfig = NN<ConsoleGetConfigQuery['core_getConfig']>;
export type IConfigs = NN<ConsoleListConfigsQuery['core_listConfigs']>;

export const iotConfigQueries = (executor: IExecutor) => ({
  updateConfig: executor(
    gql`
      mutation Core_updateConfig(
        $envName: String!
        $config: ConfigIn!
      ) {
        core_updateConfig(
          envName: $envName
          config: $config
        ) {
          id
        }
      }
    `,
    {
      transformer: (data: ConsoleUpdateConfigMutation) => data,
      vars(_: ConsoleUpdateConfigMutationVariables) {},
    }
  ),
  deleteConfig: executor(
    gql`
      mutation Core_deleteConfig(
        $envName: String!
        $configName: String!
      ) {
        core_deleteConfig(
          envName: $envName
          configName: $configName
        )
      }
    `,
    {
      transformer: (data: ConsoleDeleteConfigMutation) => data,
      vars(_: ConsoleDeleteConfigMutationVariables) {},
    }
  ),
  getConfig: executor(
    gql`
      query Core_getConfig(
        $envName: String!
        $name: String!
      ) {
        core_getConfig(
          envName: $envName
          name: $name
        ) {
          binaryData
          data
          displayName
          environmentName
          immutable
          metadata {
            annotations
            creationTimestamp
            deletionTimestamp
            generation
            labels
            name
            namespace
          }
        }
      }
    `,
    {
      transformer: (data: ConsoleGetConfigQuery) => data.core_getConfig,
      vars(_: ConsoleGetConfigQueryVariables) {},
    }
  ),
  listConfigs: executor(
    gql`
      query Core_listConfigs(
        $envName: String!
        $search: SearchConfigs
        $pq: CursorPaginationIn
      ) {
        core_listConfigs(
          envName: $envName
          search: $search
          pq: $pq
        ) {
          edges {
            cursor
            node {
              createdBy {
                userEmail
                userId
                userName
              }
              creationTime
              displayName
              data
              environmentName
              immutable
              lastUpdatedBy {
                userEmail
                userId
                userName
              }
              markedForDeletion
              metadata {
                annotations
                creationTimestamp
                deletionTimestamp
                generation
                labels
                name
                namespace
              }
              updateTime
            }
          }
          pageInfo {
            endCursor
            hasNextPage
            hasPrevPage
            startCursor
          }
          totalCount
        }
      }
    `,
    {
      transformer: (data: ConsoleListConfigsQuery) => data.core_listConfigs,
      vars(_: ConsoleListConfigsQueryVariables) {},
    }
  ),
  createConfig: executor(
    gql`
      mutation Core_createConfig(
        $envName: String!
        $config: ConfigIn!
      ) {
        core_createConfig(
          envName: $envName
          config: $config
        ) {
          id
        }
      }
    `,
    {
      transformer: (data: ConsoleCreateConfigMutation) =>
        data.core_createConfig,
      vars(_: ConsoleCreateConfigMutationVariables) {},
    }
  ),
});
