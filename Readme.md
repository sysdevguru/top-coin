# Top Coins
Simple Golang application shows top ranked cryptocurrency with prices.

## Problem & solutions
I am asked to show only top 200 cryptocurrencies only in case limit is not mentioned in the request.  
### Problems
#### Architecture
It should be fully Microservices and will be good in case it can be event-driven.  
So that any market value changes can trigger our Ranking service and Pricing service which we can deploy these services into AWS lambda.  
But it will not be available to get real time market information with free API subscription.  
#### Non consistency
CryptoCompare does not provide consistent ranking information, in some cases, some rankings are not available in the response.
#### Incompatibility
CoinMarketCap and CryptoCompare don't have exact same coin informations, for example, YBC information is not available on CoinMarketCap.  
This can make the empty prices in the JSON response.  

Note: Free subscription of CoinMarketCap API limits daily requests. So it will not be able to fetch after a limited amount of request.  

### Implementation
#### Microservices
- Web service  
- Ranking service  
- Pricing service  
- PostgreSQL database  
#### Additional parameter
Since we can have empty prices values because of the two service providers incompatibility, we add one more field `top` which is boolean to get fully available top coin informations.
```sh
curl 'localhost:8080/api/v1/coins/list?limit=50'
```
The above request will get 50 coins informations from ranking 1 to 50.  
```sh
curl 'localhost:8080/api/v1/coins/list?limit=50&top=true'
```
If we specify the `top` parameter, it will fetch top 50 fully available coin informations.  
#### IPC (Inter-Process Communication)
All the services share the PostgreSQL database in this implementation.  
This implementation is not event-driven, means it frequentely access to service providers to get updated informations.  
In this case, not shared DB but IPCs such as gRPC, socket, Pubsub can not be positive solutions.  
If we implement Pubsub, making Ranking service and Pricing service as publisher and http server as a subscriber, http server has to store the received data into DB.  
So that user can access to data at anytime.  
Publishers can store the information into DB but that will be duplicated work regarding the DB operations.  
#### Fetching market informations
Ranking service and Pricing service make requests to each information provider every 1 minute.  
They fetch ranking and pricing informations from its providers and stores informations into `price_info` and `rank_info` tables.  

## Prerequisites
### Golang, Docker, Docker Compose
https://golang.org/doc/install  
https://docs.docker.com/install  
https://docs.docker.com/compose/install  

## How to run
```sh
go get github.com/sysdevguru/top-coin
cd $GOPATH/src/github.com/sysdevguru/top-coin
docker-compose up --remove-orphans
```

From other Termial  
To get from ranking 1 to 10.  
```sh
curl 'localhost:8080/api/v1/coins/list?limit=10'
```

All Top 200 currencies.  
```sh
curl 'localhost:8080/api/v1/coins/list'
```

To get Top 20 fully available coin informations.  
```sh
curl 'localhost:8080/api/v1/coins/list?limit=20?top=true'
```

## What I want to do more or in other way
### Docker images
I would like to make each service as docker image and commit to dockerhub.  
### Test
I would like to add unit_test and integration_test using testify.  
### Auth mdw
I did not implement Authentication mdw because of the time and just let it forward all requests without checking authentication.  
In the Auth mdw, we can check authentication and do rate limiting, and upload size limiting etc  
### API Keys, credentials, configuration informations
I would like to gather all statically mentioned values into one `config.yml`  
In this yaml file, we can store postgres credential, API key, Check interval etc.  
### Http Response
I can add more details in the http response regarding the error case.  
### Daemonize
I would likt to daemonize the Pricing service [price-srv] and Ranking service [rank-srv].