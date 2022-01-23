# Hub-of-Hubs-sync-service

[![Go Report Card](https://goreportcard.com/badge/github.com/stolostron/hub-of-hubs-sync-service)](https://goreportcard.com/report/github.com/stolostron/hub-of-hubs-sync-service)
[![License](https://img.shields.io/github/license/stolostron/hub-of-hubs-sync-service)](/LICENSE)

This repository contains instructions for how to deploy 
[OpenHorizon edge sync service](https://github.com/open-horizon/edge-sync-service) 
components as part of the Hub-of-Hubs PoC.  
OpenHorizon edge sync service is used as transport layer in the PoC and have two main components:  
1.  Cloud Sync Service (CSS) running in the cloud. In this PoC the CSS runs in the hub of hubs.
1.  Edge Sync Service (ESS) running in edge nodes. In this PoC the ESS runs in leaf hubs. 

Go to the [Contributing guide](CONTRIBUTING.md) to learn how to get involved.

## Getting Started

### Cloud Sync Service (CSS)

#### Deploy CSS on Hub of Hubs cluster

1.  Set the `CSS_PORT` environment variable to hold the CSS port to be used internally in the cluster.
    ```
    $ export CSS_PORT=...
    ```
    
1.  Run the following command to deploy the CSS to your Hub of Hubs cluster:  
    ```
    envsubst < css/css.yaml.template | kubectl apply -f -
    ```
    
edge-sync-service CSS k8s objects will be created under the namespace `sync-service`.

#### Cleanup of CSS from Hub of Hubs cluster
    
1.  Run the following command to clean the CSS from your Hub of Hubs cluster:  
    ```
    envsubst < css/css.yaml.template | kubectl delete -f -
    ``` 

### Edge Sync Service (ESS)

#### Deploy ESS on a leaf hub cluster

1.  Set the `CSS_HOST` environment variable to hold the CSS host.
    ```
    $ export CSS_HOST=...
    ```
    
1.  Set the `CSS_PORT` environment variable to hold the CSS port.
    ```
    $ export CSS_PORT=...
    ```
    
1.  Set the `LISTENING_PORT` environment variable to hold the ESS http listening port.
    ```
    $ export LISTENING_PORT=...
    ```
    
1.  Set the `LH_ID` environment variable to hold a unique leaf-hub id.
    ```
    $ export LH_ID=...
    ```
    
1.  Run the following command to deploy the ESS to your leaf hub cluster:  
    ```
    envsubst < ess/ess.yaml.template | kubectl apply -f -
    ```
    
edge-sync-service ESS k8s objects will be created under the namespace `sync-service`.
    
#### Cleanup of ESS from a leaf hub cluster
    
1.  Run the following command to clean the ESS from your leaf hub cluster:  
    ```
    envsubst < ess/ess.yaml.template | kubectl delete -f -
    ``` 
