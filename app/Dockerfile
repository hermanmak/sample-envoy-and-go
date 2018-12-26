FROM golang:alpine

ADD ./src /go/src/app

WORKDIR /go/src/app

RUN go build -o myapp 

ENTRYPOINT ./myapp
