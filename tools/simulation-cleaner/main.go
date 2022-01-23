package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/open-horizon/edge-sync-service-client/client"
	datatypes "github.com/stolostron/hub-of-hubs-data-types"
)

const (
	envSyncServiceHost        = "SYNC_SERVICE_HOST"
	envSyncServicePort        = "SYNC_SERVICE_PORT"
	envLeafHubID              = "LH_ID"
	envNumOfSimulatedLeafHubs = "NUMBER_OF_SIMULATED_LEAF_HUBS"

	serverProtocol = "http"
	orgID          = "myorg"
	appKey         = "user@myorg"
)

var (
	errEnvVarNotFound = errors.New("not found environment variable")
	errWrongVarType   = errors.New("environment variable is not an integer")
)

// cleans SynService storage.
func main() {
	host, port, leafHubName, numOfSimulatedLeafHubs, err := readEnvVars()
	if err != nil {
		log.Fatalln(fmt.Sprintf("initialization error - %v", err))
	}

	syncServiceClient := client.NewSyncServiceClient(serverProtocol, host, port)

	syncServiceClient.SetOrgID(orgID)
	syncServiceClient.SetAppKeyAndSecret(appKey, "")

	cleanObjects(syncServiceClient, leafHubName)

	for i := 1; i <= numOfSimulatedLeafHubs; i++ {
		cleanObjects(syncServiceClient, fmt.Sprintf("%s_simulated_%d", leafHubName, i))
	}
}

func readEnvVars() (string, uint16, string, int, error) {
	host, found := os.LookupEnv(envSyncServiceHost)
	if !found {
		return "", 0, "", 0, fmt.Errorf("%w: %s", errEnvVarNotFound, envSyncServiceHost)
	}

	portStr, found := os.LookupEnv(envSyncServicePort)
	if !found {
		return "", 0, "", 0, fmt.Errorf("%w: %s", errEnvVarNotFound, envSyncServicePort)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return "", 0, "", 0, fmt.Errorf("%w: %s", errWrongVarType, envSyncServicePort)
	}

	leafHubName, found := os.LookupEnv(envLeafHubID)
	if !found {
		return "", 0, "", 0, fmt.Errorf("%w: %s", errEnvVarNotFound, envLeafHubID)
	}

	numOfSimulatedLeafHubsStr, found := os.LookupEnv(envNumOfSimulatedLeafHubs)
	if !found {
		return "", 0, "", 0, fmt.Errorf("%w: %s", errEnvVarNotFound, envNumOfSimulatedLeafHubs)
	}

	numOfSimulatedLeafHubs, err := strconv.Atoi(numOfSimulatedLeafHubsStr)
	if err != nil {
		return "", 0, "", 0, fmt.Errorf("%w: %s", errWrongVarType, envNumOfSimulatedLeafHubs)
	}

	return host, uint16(port), leafHubName, numOfSimulatedLeafHubs, nil
}

func cleanObjects(client *client.SyncServiceClient, leafHubName string) {
	log.Println(fmt.Sprintf("cleaning old data of leaf hub - %s", leafHubName))

	objectIDs := []string{
		fmt.Sprintf("%s.%s", leafHubName, datatypes.ManagedClustersMsgKey),
		fmt.Sprintf("%s.%s", leafHubName, datatypes.ClustersPerPolicyMsgKey),
		fmt.Sprintf("%s.%s", leafHubName, datatypes.PolicyCompleteComplianceMsgKey),
		fmt.Sprintf("%s.%s", leafHubName, datatypes.PolicyDeltaComplianceMsgKey),
		fmt.Sprintf("%s.%s", leafHubName, datatypes.MinimalPolicyComplianceMsgKey),
		fmt.Sprintf("%s.%s", leafHubName, datatypes.ControlInfoMsgKey),
		fmt.Sprintf("%s.%s", leafHubName, datatypes.LocalPlacementRulesMsgKey),
		fmt.Sprintf("%s.%s", leafHubName, datatypes.LocalClustersPerPolicyMsgKey),
		fmt.Sprintf("%s.%s", leafHubName, datatypes.LocalPolicyCompleteComplianceMsgKey),
		fmt.Sprintf("%s.%s", leafHubName, datatypes.LocalPolicySpecMsgKey),
	}

	for _, objectID := range objectIDs {
		if err := client.DeleteObject(datatypes.StatusBundle, objectID); err != nil {
			log.Println(fmt.Sprintf("failed to clean old data for %s, %v", objectID, err))
		} else {
			log.Println(fmt.Sprintf("successfully cleaned old data for %s", objectID))
		}
	}
}
