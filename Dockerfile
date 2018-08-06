FROM golang:1.9-alpine AS build-dev
RUN apk update
RUN apk add --no-cache \
      ca-certificates \
      openssl \
      git \
      make \
      bash \
      gcc \
      musl-dev

COPY . /go/src/github.com/currencies-exchange
WORKDIR /go/src/github.com/currencies-exchange

# Pull required packages
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

RUN go build
CMD ["/go/src/github.com/currencies-exchange/entrypoint.sh"]
