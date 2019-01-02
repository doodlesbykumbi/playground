FROM golang:latest as builder

RUN go get github.com/cespare/reflex

RUN mkdir /app
WORKDIR /app

COPY . .

CMD reflex -r '\.go$' -s -- bash -c "go run main.go"
