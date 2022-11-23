# Go Kafka

Dockerfile example with specified platform:
```
FROM --platform=linux/amd64 golang:1.15 // <---- here
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
ENTRYPOINT go build -o main .
```

## Prerequisites

- Docker installed - (Tested with Docker 4.7.0 (77141) but latest should work)
- Make sure docker has a minimum of 6GB ram allocated as per Confluent Kafka docs. (Tested with 8GB ram, 1GB Swap and 5 CPUS)

## Steps to run

Clone and navigate to the root of this repo in your terminal.

- Create docker bridge network
```
docker network create go-confluent-network
```

- Run Confluent Kafka services
```
docker-compose up
```

- Run Basic Go Kafka Producer API example
```
cd ./producer && docker-compose up
```

- Run Basic Go Kafka Consumer API example
```
cd ../consumer && docker-compose up
```
or
- Run all the above command using the run.sh script
```
sh run.sh
```
The above containers/services are setup on same brige network for them to be able to communicate with each other.<br />
You will also see that the Consumer/Producer DockerFile EntryPoint has CompileDaemon within it, which watches your .go files and builds and runs on file change for fast developement in Docker.

## Consumer and Producer API usage

- (POST) http://localhost:8090/api/producer/:topic/:data ~ {topic} - will create topic if doesn't exist already ~ {data} - no schema set so add any data.
- (GET) http://localhost:8091/api/consumer/:topic ~ returns consumed data for the specified topic

### Confluent Control Center
- Is accessible via  ~ http://localhost:9021


## Docs

[Docker Multi Arch Images](https://www.docker.com/blog/multi-arch-images/tested)<br />
[Confluent Kafka Docker](https://docs.confluent.io/platform/current/quickstart/ce-docker-quickstart.html)<br />
[Go Compile Daemon](https://github.com/githubnemo/CompileDaemon)

