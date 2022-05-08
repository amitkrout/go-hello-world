Deploy and run a simple/basic go web app in containerized environment
---

### Overview

Aim of this repository is to start an HTTP server and serve a static web page. However you can keep adding your program output in the static web page and can be verified from your browser. This can be run directly from go code and can be built, deployed on container orchestration tool like k8s, k3s and OpenShift.

### How to run locally using go run time

```shell
$ git clone https://github.com/amitkrout/go-hello-world
$ cd go-hello-world
$ go run main.go
Server running (port=8080), route: http://localhost:8080/helloworld
```
Access the URL from your browser - http://localhost:8080/helloworld. It should display `Hello, World!`

### How to run locally using docker run time

```shell
$ git clone https://github.com/amitkrout/go-hello-world
$ cd go-hello-world
$ docker build -t hello_go_http .
$ docker run -p 8080:8080 -t hello_go_http
Server running (port=8080), route: http://localhost:8080/helloworld
```

Access the URL from your browser - http://localhost:8080/helloworld. It should display `Hello, World!`


### How to run locally go app server in k3s

```shell
$ k3d registry create myregistry.localhost --port 12345
$ k3d cluster create --servers 3 --image rancher/k3s:v1.22.5-k3s1 --registry-use k3d-myregistry.localhost:12345 --api-port 6550 -p "8081:80@loadbalancer" --agents 2
$ docker build . --file Dockerfile --tag k3d-myregistry.localhost:12345/hello-go-http:v1.0
$ docker push k3d-myregistry.localhost:12345/hello-go-http:v1.0
$ kubectl create deployment web --image=k3d-myregistry.localhost:12345/hello-go-http:v1.0
$ kubectl create service clusterip web --tcp=8080:8080
$ cat <<EOF | kubectl apply -f -
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: web
  annotations:
    ingress.kubernetes.io/ssl-redirect: "false"
spec:
  rules:
  - http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: web
            port:
              number: 8080
EOF
$ curl http://localhost:8081/helloworld
```

### How to run integration and e2e locally

```shell
$ go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo && export PATH=$PATH:$(go env GOPATH)/bin/
$ git clone https://github.com/amitkrout/go-hello-world
$ cd go-hello-world
$ ginkgo -focus="Feature integration test" tests/integration/
$ ginkgo -focus="End to end scenario" tests/e2e/
```
NOTE: GinkGo is on-boarded as the test framework to write integration and end to end test. Please visit [GinkGo](https://onsi.github.io/ginkgo/) documentation to know more about and the features it has.
