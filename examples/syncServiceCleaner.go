package examples

import (
	"fmt"
	"github.com/open-horizon/edge-sync-service-client/client"
	"os"
	"strconv"
)

const (
	envSyncServiceHost          = "SYNC_SERVICE_HOST"
	envSyncServicePort          = "SYNC_SERVICE_PORT"
	envSyncServiceOrgId         = "SYNC_SERVICE_ORG_ID"
	envSyncServiceAppKey        = "SYNC_SERVICE_APP_KEY"
	envSyncServiceLeafHubName   = "SYNC_SERVICE_LEAF_HUB_NAME"
	envSyncServiceNumOfLeafHubs = "SYNC_SERVICE_NUM_OF_LEAF_HUBS"

	serverProtocol = "http"

	StatusBundle                  = "StatusBundle"
	ManagedClustersMsgKey         = "ManagedClusters"
	ClustersPerPolicyMsgKey       = "ClustersPerPolicy"
	PolicyComplianceMsgKey        = "PolicyCompliance"
	MinimalPolicyComplianceMsgKey = "MinimalPolicyCompliance"
	ControlInfoMsgKey             = "ControlInfo"
)

// Clean cleans SynService storage.
func Clean() {
	host, found := os.LookupEnv(envSyncServiceHost)
	if !found {
		fmt.Printf("Environment variable %s not found", envSyncServiceHost)
		return
	}

	portStr, found := os.LookupEnv(envSyncServicePort)
	if !found {
		fmt.Printf("Environment variable %s not found", envSyncServicePort)
		return
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("Environment variable %s is not an integer", envSyncServicePort)
		return
	}

	orgId, found := os.LookupEnv(envSyncServiceOrgId)
	if !found {
		fmt.Printf("Environment variable %s not found", envSyncServiceOrgId)
		return
	}

	appKey, found := os.LookupEnv(envSyncServiceAppKey)
	if !found {
		fmt.Printf("Environment variable %s not found", envSyncServiceAppKey)
		return
	}

	leafHubName, found := os.LookupEnv(envSyncServiceLeafHubName)
	if !found {
		fmt.Printf("Environment variable %s not found", envSyncServiceLeafHubName)
		return
	}

	numOfLeafHubsStr, found := os.LookupEnv(envSyncServiceNumOfLeafHubs)
	if !found {
		fmt.Printf("Environment variable %s not found", envSyncServiceNumOfLeafHubs)
		return
	}

	numOfLeafHubs, err := strconv.Atoi(numOfLeafHubsStr)
	if err != nil {
		fmt.Printf("Environment variable %s is not an integer", envSyncServiceNumOfLeafHubs)
		return
	}

	syncServiceClient := client.NewSyncServiceClient(serverProtocol, host, uint16(port))

	syncServiceClient.SetOrgID(orgId)
	syncServiceClient.SetAppKeyAndSecret(appKey, "")

	fmt.Printf("cleaning old data\n")

	clean(syncServiceClient, leafHubName)

	for i := 0; i <= numOfLeafHubs; i++ {
		clean(syncServiceClient, fmt.Sprintf("%s_simulated_%d", leafHubName, i))
	}
}

func clean(client *client.SyncServiceClient, leafHubName string) {
	objectIDs := make([]string, 0)

	objectIDs = append(objectIDs, fmt.Sprintf("%s.%s", leafHubName, ManagedClustersMsgKey))
	objectIDs = append(objectIDs, fmt.Sprintf("%s.%s", leafHubName, ClustersPerPolicyMsgKey))
	objectIDs = append(objectIDs, fmt.Sprintf("%s.%s", leafHubName, PolicyComplianceMsgKey))
	objectIDs = append(objectIDs, fmt.Sprintf("%s.%s", leafHubName, MinimalPolicyComplianceMsgKey))
	objectIDs = append(objectIDs, fmt.Sprintf("%s.%s", leafHubName, ControlInfoMsgKey))

	for _, objectID := range objectIDs {
		if err := client.DeleteObject(StatusBundle, objectID); err != nil {
			fmt.Printf("failed to clean old data for %s, %s, %v\n", leafHubName, objectID, err)
		} else {
			fmt.Printf("successfuly cleaned old data for %s, %s\n", leafHubName, objectID)
		}
	}
}
