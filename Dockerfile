FROM golang:1.17-alpine

RUN apk update && apk add curl \
                          git \
                          protobuf \
                          bash \
                          make \
                          openssh-client && \
     rm -rf /var/cache/apk/*

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN make setup

RUN make run


EXPOSE 8080
