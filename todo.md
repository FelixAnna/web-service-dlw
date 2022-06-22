## TODO list
### research on DI 
    -- wire : compile-time dependency injection  (done)
    -- dig /inject: runtime dependency injection
### research helm upgrade not update replicaset + pod when image version changed
### delete stucked namespace
    https://phoenixnap.com/kb/kubernetes-delete-namespace
### ut
    ```
    go test -cover -v ./...
    ```
### document api

### ut with mockery
mockery --all --recursive --output=./mocks --with-expecter