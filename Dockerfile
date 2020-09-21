FROM golang:1.15.2
 
WORKDIR /go/src/cepservice
 
ADD . /go/src/cepservice

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

RUN go build

EXPOSE 8000

CMD ["./cepservice"]