#!/usr/bin/env bash


docker run -ti --rm -p 8080:8080 soloio/petstore-example:latest


/Users/ilackarms/go/src/github.com/solo-io/gloo/_output/glooctl-darwin-amd64 add route \
    --path-exact /sample-route-1 \
    --dest-name default-petstore-8080 \
    --prefix-rewrite /api/pets --yaml > vs.yaml

/Users/ilackarms/go/src/github.com/solo-io/gloo/_output/glooctl-darwin-amd64 create us static \
    --name default-petstore-8080 \
    --static-hosts 10.1.10.116:8080 --yaml > petstore-upstream.yaml

consul kv put gloo/gloo.solo.io/v1/Upstream/gloo-system @petstore-upstream.yaml
consul kv put gloo/gateway.solo.io/v1/VirtualService/gloo-system @vs.yaml
