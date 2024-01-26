# POC Redis Distributed Lock

## How to run
```sh
make new-network # for new docker-network
make some-redis # start redis
docker-compose up -d --scale myapp=3 # run-api 3 instance
```