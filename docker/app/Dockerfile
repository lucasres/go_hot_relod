FROM golang:1.16-alpine

RUN go install github.com/canthefason/go-watcher/cmd/watcher@latest

WORKDIR /app

COPY . .

ENTRYPOINT go run main.go