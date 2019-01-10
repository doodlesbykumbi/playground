#!/usr/bin/env bash

protoc -I ../../helloworld \
 --go_out=plugins=grpc:../../helloworld \
 --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../../client_js/src/proto \
 --js_out=import_style=commonjs:../../client_js/src/proto \
 ../../helloworld/helloworld.proto

disable_eslint() {
  echo -e "/* eslint-disable */\n$(cat $1)" > $1
}

disable_eslint ../../client_js/src/proto/helloworld_grpc_web_pb.js
disable_eslint ../../client_js/src/proto/helloworld_pb.js
