FROM --platform=linux/amd64 golang:1.15
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build -o main ." --command=/app/main