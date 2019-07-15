FROM golang:1.12


COPY . /go/src/semantic-repository
WORKDIR /go/src/semantic-repository

ENV GO111MODULE=on

RUN go build

EXPOSE 8080

CMD ./semantic-repository