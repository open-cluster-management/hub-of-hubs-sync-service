# simulation-cleaner

This package contains instructions how to build and run sync service simulation storage cleaner 

1. Build the code by running the following command: `make`

1. Set the `SYNC_SERVICE_HOST` environment variable to hold the CSS host.
    ```
    $ export SYNC_SERVICE_HOST=...
    ```
    
1. Set the `SYNC_SERVICE_PORT` environment variable to hold the CSS port.
    ```
    $ export SYNC_SERVICE_PORT=...
    ```
    
1. Set the `LH_ID` environment variable to hold the leaf hub unique id as used in the simulation.
    ```
    $ export LH_ID=...
    ```

1. Set the `NUMBER_OF_SIMULATED_LEAF_HUBS` environment variable to hold the number of simulated leaf hubs (not including the original leaf hub).
    ```
    $ export NUMBER_OF_SIMULATED_LEAF_HUBS=...
    ```  

1. Run the following command to clean sync service storage:  
    ```
    ./bin/simulation-cleaner
    ```
