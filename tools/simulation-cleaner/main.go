package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/open-horizon/edge-sync-service-client/client"
)

const (
	envSyncServiceHost          = "SYNC_SERVICE_HOST"
	envSyncServicePort          = "SYNC_SERVICE_PORT"
	envSyncServiceOrgID         = "SYNC_SERVICE_ORG_ID"
	envSyncServiceAppKey        = "SYNC_SERVICE_APP_KEY"
	envSyncServiceLeafHubName   = "SYNC_SERVICE_LEAF_HUB_NAME"
	envSyncServiceNumOfLeafHubs = "SYNC_SERVICE_NUM_OF_LEAF_HUBS"

	serverProtocol = "http"

	statusBundle                  = "StatusBundle"
	managedClustersMsgKey         = "ManagedClusters"
	clustersPerPolicyMsgKey       = "ClustersPerPolicy"
	policyComplianceMsgKey        = "PolicyCompliance"
	minimalPolicyComplianceMsgKey = "MinimalPolicyCompliance"
	controlInfoMsgKey             = "ControlInfo"
)

// clean cleans SynService storage.
func clean() {
	host, found := os.LookupEnv(envSyncServiceHost)
	if !found {
		fmt.Printf("Environment variable %s not found\n", envSyncServiceHost)
		return
	}

	portStr, found := os.LookupEnv(envSyncServicePort)
	if !found {
		fmt.Printf("Environment variable %s not found\n", envSyncServicePort)
		return
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("Environment variable %s is not an integer\n", envSyncServicePort)
		return
	}

	orgID, found := os.LookupEnv(envSyncServiceOrgID)
	if !found {
		fmt.Printf("Environment variable %s not found\n", envSyncServiceOrgID)
		return
	}

	appKey, found := os.LookupEnv(envSyncServiceAppKey)
	if !found {
		fmt.Printf("Environment variable %s not found\n", envSyncServiceAppKey)
		return
	}

	leafHubName, found := os.LookupEnv(envSyncServiceLeafHubName)
	if !found {
		fmt.Printf("Environment variable %s not found\n", envSyncServiceLeafHubName)
		return
	}

	numOfLeafHubsStr, found := os.LookupEnv(envSyncServiceNumOfLeafHubs)
	if !found {
		fmt.Printf("Environment variable %s not found\n", envSyncServiceNumOfLeafHubs)
		return
	}

	numOfLeafHubs, err := strconv.Atoi(numOfLeafHubsStr)
	if err != nil {
		fmt.Printf("Environment variable %s is not an integer\n", envSyncServiceNumOfLeafHubs)
		return
	}

	syncServiceClient := client.NewSyncServiceClient(serverProtocol, host, uint16(port))

	syncServiceClient.SetOrgID(orgID)
	syncServiceClient.SetAppKeyAndSecret(appKey, "")

	fmt.Printf("cleaning old data\n")

	cleanObjects(syncServiceClient, leafHubName)

	for i := 0; i <= numOfLeafHubs; i++ {
		cleanObjects(syncServiceClient, fmt.Sprintf("%s_simulated_%d", leafHubName, i))
	}
}

func cleanObjects(client *client.SyncServiceClient, leafHubName string) {
	objectIDs := make([]string, 0)

	objectIDs = append(objectIDs, fmt.Sprintf("%s.%s", leafHubName, managedClustersMsgKey))
	objectIDs = append(objectIDs, fmt.Sprintf("%s.%s", leafHubName, clustersPerPolicyMsgKey))
	objectIDs = append(objectIDs, fmt.Sprintf("%s.%s", leafHubName, policyComplianceMsgKey))
	objectIDs = append(objectIDs, fmt.Sprintf("%s.%s", leafHubName, minimalPolicyComplianceMsgKey))
	objectIDs = append(objectIDs, fmt.Sprintf("%s.%s", leafHubName, controlInfoMsgKey))

	for _, objectID := range objectIDs {
		if err := client.DeleteObject(statusBundle, objectID); err != nil {
			fmt.Printf("failed to clean old data for %s, %s, %v\n", leafHubName, objectID, err)
		} else {
			fmt.Printf("successfuly cleaned old data for %s, %s\n", leafHubName, objectID)
		}
	}
}

func main() {
	clean()
}
