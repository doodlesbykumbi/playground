#!/usr/bin/env bash

ip=$(minikube ip):$(kubectl get svc knative-ingressgateway --namespace istio-system --output 'jsonpath={.spec.ports[?(@.port==80)].nodePort}')

curl -v -H "Host: helloworld-go.default.example.com" "http://$ip?sleep=3000&prime=10000"

echo "
GET http://$ip?sleep=3000&prime=10000
Host: helloworld-go.default.example.com
" | vegeta attack -duration=100s -rate=100 -workers=200

