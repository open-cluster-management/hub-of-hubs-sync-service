# Hub-of-Hubs-sync-service examples

This package contains instructions how to build and run sync service storage cleaner 

1. Build the example by running the following command: `make`

1. Set the `SYNC_SERVICE_HOST` environment variable to hold the CSS host.
    ```
    $ export SYNC_SERVICE_HOST=...
    ```
    
1. Set the `SYNC_SERVICE_PORT` environment variable to hold the CSS port.
    ```
    $ export SYNC_SERVICE_PORT=...
    ```
    
1. Set the `SYNC_SERVICE_ORG_ID` environment variable to hold the CSS organization id.
    ```
    $ export SYNC_SERVICE_ORG_ID=...
    ```

1. Set the `SYNC_SERVICE_APP_KEY` environment variable to hold the CSS application key.
    ```
    $ export SYNC_SERVICE_APP_KEY=...
    ```
    
1. Set the `SYNC_SERVICE_LEAF_HUB_NAME` environment variable to hold the name of the leaf hub used in the simulation.
    ```
    $ export SYNC_SERVICE_LEAF_HUB_NAME=...
    ```

1. Set the `SYNC_SERVICE_NUM_OF_LEAF_HUBS` environment variable to hold the total number of leaf hubs used in the simulation.
    ```
    $ export SYNC_SERVICE_NUM_OF_LEAF_HUBS=...
    ```  

1. Run the following command to clean sync service storage:  
    ```
    ./bin/hub-of-hubs-sync-service
    ```
