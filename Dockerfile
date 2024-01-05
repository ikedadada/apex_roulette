FROM golang:1.21.5-alpine AS builder

ENV GOPATH /go

RUN mkdir -p /go/src
WORKDIR /go/src

COPY . /go/src

RUN go mod tidy && \
    go build -o /go/bin/app

CMD ["/go/bin/app"]
