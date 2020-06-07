# Stocks-GraphQL
GraphQL API to retrieve Historical Stocks data from DynamoDB

### Features:

* Implements Querys, Mutations and Subscriptions (https://medium.com/software-insight/graphql-queries-mutations-and-subscriptions-286522b263d9)
* Data effeciently indexed and stored on DynamoDB 
* Database contains 9 Million records of historical data
* Project deployed on Google App Engine
* Implements OAuth 2.0 with Google Integration
* **Automatic Persisted Query (APQ)** Redis Cache hosted on Redis Labs Enterprise Free Tier

``` diff 
! API can only be accessed on the GraphQL playground GUI
```

### Environment Variables

`PORT` // PORT to run your application </br>
`AWS_ACCESS_TOKEN` // Your AWS Access Token </br>
`AWS_SECRET` // your AWS Secret </br>
`REDIS_ADDRESS` _<host_name>:<port_name>_

## To Run the project:

* Install and setup Go 
    - https://golang.org/doc/install
    - Go Modules: https://github.com/golang/go/wiki/Modules
* `go mod download` to download the requirements
* `go mod vendor` to download to local reposistory
* Start the application:  `go run main.go`</br>


**Local Endpoint:** `localhost:8080`

## To deploy on App Engine:

**Follow Instructions from official website:** https://cloud.google.com/appengine/docs/standard/go/building-app

#### app.yaml:

```
runtime: go113

env_variables:
  PORT: <port>
  CALLBACK_URL: https://graphql-project-278000.ue.r.appspot.com/callback
  GOOGLE_CLIENT_ID: <google_client_id>
  GOOGLE_CLIENT_SECRET:<google_client_secret>
  AWS_ACCESS_TOKEN:  <aws_access_token>
  AWS_SECRET:  <aws_secret>
  REDIS_ADDRESS:  <redis_address_with_port>
  REDIS_PASS:  <redis_password>
```
