## TODO list
### research on DI 
    -- wire : compile-time dependency injection  (done)
    -- dig /inject: runtime dependency injection


### delete stucked namespace (done)
    https://phoenixnap.com/kb/kubernetes-delete-namespace
    
### ut   (done)
    ```
    go test -cover -v ./...
    ```
    
### ut with mockery   (done)
mockery --all --recursive --output=./mocks --with-expecter
mockery --name=ProblemService --recursive --output=./mocks --with-expecter

### document api    

### research helm upgrade not update replicaset + pod when image version changed
