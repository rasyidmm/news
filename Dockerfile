FROM golang:1.18-alpine3.16

RUN apk update && apk upgrade && apk add curl \
                          git \
                          protobuf \
                          bash \
                          make \
                         busybox-extras \
                        openssh-client && \
     rm -rf /var/cache/apk/*

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN make build

EXPOSE 8080
#RUN make rest
ENTRYPOINT ["./main"]
