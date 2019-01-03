#!/usr/bin/env bash

docker build -t codebykumbi/helloworld-go .
docker push codebykumbi/helloworld-go
kubectl apply --filename service.yaml
