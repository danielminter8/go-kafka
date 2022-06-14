# Working with Go Confluent Kafka library on Apple Silicone

## This is a step by step approach to running, developing and working with the Golang Confluent Kafka library on an ARM based device/Apple Silicone(M1) using just Docker.
<br />

### Some context:

Confluent Kafka has dependencies that currently does not support based devices, therefore not allowing you run Confluent Kafka. This guides shows you a possible solution to being able to run and develope on your arm device.

### How is this possible?

This approach takes advantage of Docker multi-arch builds or also know as Docker Buildx and Docker Desktop emulation. Currently if you build an docker image, docker builds the image for the host platform. So if you building the image on a M1 powered macbook, it would build an image for arm64 platform so that it would be able to run on the device it is being built on without any problems. 

So the solution to being able to run an amd64 dependent libraries in a Go APP/API on a arm based system is to specify the platform in the DockerFile and build it, thanks to there multi-arch image support that allows us to easier 'develop containers on, and for Arm servers and devices' and run that image through Docker Desktop emulation.
<br />
<br />


Dockerfile example with specified platform:
```
FROM --platform=linux/amd64 golang:1.15 // <---- here
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
ENTRYPOINT go build -o main .
```

Note: This isn't guaranteed to be stable, but it worked out for me : )

## Prerequisites

- Docker installed - (Tested with Docker 4.7.0 (77141) but latest should work)
- Make sure docker has a minimum of 6GB ram allocated as per Confluent Kafka docs. (Tested with 8GB ram, 1GB Swap and 5 CPUS)

## Steps to run

Clone and navigate to the root of this repo in your terminal.

- Run Confluent Kafka services
```
docker-compose up -d
```

- Run Basic Go Kafka Producer API example
```
cd ./producer && docker-compose up -d
```

- Run Basic Go Kafka Consumer API example
```
cd ../consumer && docker-compose up -d
```
or
- Run all the above via the run.sh script
```
sh run.sh
```
The above containers/services are setup on same brige network for them to be able to communicate with each other.<br />
You will also see that the Consumer/Producer DockerFile EntryPoint has CompileDaemon within it, which watches your .go files and builds and runs on file change for fast developement in Docker.

## Consumer and Producer API usage

- (POST) http://localhost:8090/api/producer/:topic/:data ~ {topic} - will create topic if doesn't exist already ~ {data} - no schema set so add any data.
- (GET) http://localhost:8091/api/consumer/:topic ~ using sockets to live stream kafka data straight client as soon as it comes in (so that you can see it as soon as its produced/sent)

### Confluent Control Center
- Is accessible via  ~ http://localhost:9021


## Docs

[Docker Multi Arch Images](https://www.docker.com/blog/multi-arch-images/tested)<br />
[Confluent Kafka Docker](https://docs.confluent.io/platform/current/quickstart/ce-docker-quickstart.html)<br />
[Go Compile Daemon](https://github.com/githubnemo/CompileDaemon)

## More

This README is not perfect, so if you find a mistake don't hesitate to make changes and create a pull request