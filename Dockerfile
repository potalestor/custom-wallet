FROM golang:1.15.6-alpine3.12 as builder

COPY . /go/src/github.com/potalestor/custom-wallet
WORKDIR /go/src/github.com/potalestor/custom-wallet
RUN go mod download
WORKDIR /go/src/github.com/potalestor/custom-wallet/cmd/custom-wallet

RUN go build 

EXPOSE 8080:8080

ENTRYPOINT /go/src/github.com/potalestor/custom-wallet/cmd/custom-wallet/custom-wallet --host=postgres
