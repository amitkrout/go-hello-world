# go-hello-world

## How to run locally go app server using docker

$ git clone https://github.com/amitkrout/go-hello-world

$ cd go-hello-world

$ docker build -t hello_go_http .

$ docker run -p 8080:8080 -t hello_go_http

Server running (port=8080), route: http://localhost:8080/helloworld

Access the URL from your browser - http://localhost:8080/helloworld. It should display `Hello, World!`


## How to run locally go app server in k8s

// TODO

## How to run go app server in CI

// TODO

## How to test ginkgo locally

$ go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo && export PATH=$PATH:$(go env GOPATH)/bin/

$ git clone https://github.com/amitkrout/go-hello-world

$ cd go-hello-world

$ ginkgo -focus="Feature-1 integration test" tests/integration/

$ ginkgo -focus="End to end scenario - 1" tests/e2e/