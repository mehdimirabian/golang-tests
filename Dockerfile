# Build stage
FROM golang:1.16 AS build-env

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /uds

COPY . .
RUN go build .

# Final stage
FROM ubuntu:18.04

RUN apt-get update
RUN apt-get install -y socat
VOLUME /tmp
CMD socat UNIX-LISTEN:/tmp/go.sock -

COPY --from=build-env /uds .
CMD ["./golang-tests.com"]