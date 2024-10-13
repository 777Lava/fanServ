FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN apk update && \
    apk add build-base

RUN go mod download

COPY ./ ./

CMD ["go", "run", "cmd/main.go"]
