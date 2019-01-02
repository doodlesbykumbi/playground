FROM golang:latest as builder

RUN go get github.com/cespare/reflex
RUN go get github.com/golang/mock/gomock
RUN go install github.com/golang/mock/mockgen

RUN mkdir /app
WORKDIR /app

ADD go.mod /app
ADD go.sum /app

COPY . .

RUN go mod download
RUN go generate ./pkg

CMD reflex -r '\.go$' -s -- bash -c "go run cmd/main.go"
