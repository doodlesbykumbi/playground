FROM golang:latest
ENV PROTOC_ZIP=protoc-3.3.0-linux-x86_64.zip
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.3.0/$PROTOC_ZIP
RUN apt-get update && apt-get install -y unzip
RUN unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
RUN rm -f $PROTOC_ZIP

RUN mkdir -p /work
WORKDIR /work
ADD go.mod /work
ADD go.sum /work
RUN go mod download
