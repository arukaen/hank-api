FROM golang:1.14-alpine

RUN apk add --no-cache build-base

ADD . /go/src/github.com/arukaen/hank-api
WORKDIR /go/src/github.com/arukaen/hank-api

RUN go build  -o /hank-api

EXPOSE 8080

ENTRYPOINT ["/hank-api", "-d"]
