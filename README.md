# Stocks-GraphQL
GraphQL API to retrieve Historical Stocks data from DynamoDB

### Environment Variables

`PORT` // PORT to run your application </br>
`AWS_ACCESS_TOKEN` // Your AWS Access Token </br>
`AWS_SECRET` // your AWS Secret </br>
`REDIS_ADDRESS` _<host_name>:<port_name>_ // Redis Cache Host and Port (6379) 

## To Run the project:

* Install and setup Go 
    - https://golang.org/doc/install
    - Go Modules: https://github.com/golang/go/wiki/Modules
* `go mod download` to download the requirements
* `go mod vendor` to download to local reposistory
* Start the application:  `go run main.go`</br>


**Endpoint:** `localhost:8080`


