FROM golang:1.8
MAINTAINER Luc CHMIELOWSKI <luc.chmielowski@gmail.com>

# Envs
ENV GO15VENDOREXPERIMENT=1

EXPOSE 5002
RUN apt-get update -y
RUN apt-get install libsasl2-dev
RUN mkdir -p /go/src/github.com/iochti/thing-service
WORKDIR /go/src/github.com/iochti/thing-service
COPY . /go/src/github.com/iochti/thing-service

RUN go get github.com/tools/godep
RUN godep restore
RUN go install ./...

RUN rm -rf /go/src/github.com/iochti/thing-service

CMD ["thing-service"]
