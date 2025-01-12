import gql from 'graphql-tag';
import { IExecutor } from '~/root/lib/server/helpers/execute-query-with-context';
import { NN } from '~/root/lib/types/common';
import {
  ConsoleCreateByokClusterMutation,
  ConsoleCreateByokClusterMutationVariables,
  ConsoleDeleteByokClusterMutation,
  ConsoleDeleteByokClusterMutationVariables,
  ConsoleGetByokClusterInstructionsQuery,
  ConsoleGetByokClusterInstructionsQueryVariables,
  ConsoleGetByokClusterQuery,
  ConsoleGetByokClusterQueryVariables,
  ConsoleListByokClustersQuery,
  ConsoleListByokClustersQueryVariables,
  ConsoleUpdateByokClusterMutation,
  ConsoleUpdateByokClusterMutationVariables,
} from '~/root/src/generated/gql/server';

export type IByocClusters = NN<
  ConsoleListByokClustersQuery['infra_listBYOKClusters']
>;
export type IByocCluster = NN<
  ConsoleGetByokClusterQuery['infra_getBYOKCluster']
>;

export const byokClusterQueries = (executor: IExecutor) => ({
  deleteByokCluster: executor(
    gql`
      mutation Infra_deleteBYOKCluster($name: String!) {
        infra_deleteBYOKCluster(name: $name)
      }
    `,
    {
      transformer: (data: ConsoleDeleteByokClusterMutation) =>
        data.infra_deleteBYOKCluster,
      vars(_: ConsoleDeleteByokClusterMutationVariables) {},
    }
  ),
  createBYOKCluster: executor(
    gql`
      mutation Infra_createBYOKCluster($cluster: BYOKClusterIn!) {
        infra_createBYOKCluster(cluster: $cluster) {
          id
        }
      }
    `,
    {
      transformer: (data: ConsoleCreateByokClusterMutation) =>
        data.infra_createBYOKCluster,
      vars(_: ConsoleCreateByokClusterMutationVariables) {},
    }
  ),
  updateByokCluster: executor(
    gql`
      mutation Infra_updateBYOKCluster(
        $clusterName: String!
        $displayName: String!
      ) {
        infra_updateBYOKCluster(
          clusterName: $clusterName
          displayName: $displayName
        ) {
          id
        }
      }
    `,
    {
      transformer: (data: ConsoleUpdateByokClusterMutation) =>
        data.infra_updateBYOKCluster,
      vars(_: ConsoleUpdateByokClusterMutationVariables) {},
    }
  ),
  getBYOKClusterInstructions: executor(
    gql`
      query Infrat_getBYOKClusterSetupInstructions(
        $name: String!
        $onlyHelmValues: Boolean
      ) {
        infrat_getBYOKClusterSetupInstructions(
          name: $name
          onlyHelmValues: $onlyHelmValues
        ) {
          command
          title
        }
      }
    `,
    {
      transformer: (data: ConsoleGetByokClusterInstructionsQuery) =>
        data.infrat_getBYOKClusterSetupInstructions,
      vars(_: ConsoleGetByokClusterInstructionsQueryVariables) {},
    }
  ),
  getByokCluster: executor(
    gql`
      query Infra_getBYOKCluster($name: String!) {
        infra_getBYOKCluster(name: $name) {
          accountName
          ownedBy
          createdBy {
            userEmail
            userId
            userName
          }
          creationTime
          displayName
          id
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
          recordVersion
          syncStatus {
            action
            error
            lastSyncedAt
            recordVersion
            state
            syncScheduledAt
          }
          updateTime
          clusterSvcCIDR
          globalVPN
        }
      }
    `,
    {
      transformer: (data: ConsoleGetByokClusterQuery) =>
        data.infra_getBYOKCluster,
      vars(_: ConsoleGetByokClusterQueryVariables) {},
    }
  ),
  listByokClusters: executor(
    gql`
      query Infra_listBYOKClusters(
        $search: SearchCluster
        $pagination: CursorPaginationIn
      ) {
        infra_listBYOKClusters(search: $search, pagination: $pagination) {
          edges {
            cursor
            node {
              accountName
              ownedBy
              clusterSvcCIDR
              lastOnlineAt
              createdBy {
                userEmail
                userId
                userName
              }
              creationTime
              displayName
              globalVPN
              id
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
              recordVersion
              syncStatus {
                action
                error
                lastSyncedAt
                recordVersion
                state
                syncScheduledAt
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
      transformer: (data: ConsoleListByokClustersQuery) => {
        return data.infra_listBYOKClusters;
      },
      vars(_: ConsoleListByokClustersQueryVariables) {},
    }
  ),
});
