# Top Coins
Simple Golang application shows top ranked cryptocurrency with prices.

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
