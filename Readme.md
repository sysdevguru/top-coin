# Top Coins
Simple Golang application shows top ranked cryptocurrency with prices.

## Problem & solutions
### Problems
Since ranking service and pricing service has to update ranking and price information real time, the higher subscription was needed for fetching their real time information.  
Using free subscription, we just have to request frequently to the server to fetch informations.  

I am asked to show only top 200 cryptocurrencies only in case limit is not mentioned in the request.  

### Implementation
Ranking service and Pricing service make requests to each information provider every 1 minute.  
They fetch ranking and pricing informations from its providers and stores informations into `price_info` and `rank_info` tables.  

## Prerequisites
### Docker, Docker Compose
https://docs.docker.com/install  
https://docs.docker.com/compose/install  

## How to run
```sh
go get github.com/sysdevguru/top-coin
cd $GOPATH/src/github.com/sysdevguru/top-coin
docker-compose up --remove-orphans
```

From other Termial or Postman  
Top 10 ranked currency.  
```sh
curl localhost:8080/api/v1/coins/list?limit=10
```

All currencies.  
```sh
curl -X POST localhost:8080/api/v1/coins/list
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
I would likt to make the Pricing service [price-srv] and Ranking service [rank-srv] as a daemon.  
