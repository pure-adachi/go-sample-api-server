FROM golang:1.17.1-alpine3.14

RUN apk update \
  && go get -u github.com/cosmtrek/air \
  && chmod +x ${GOPATH}/bin/air

WORKDIR /myapp
