# go-hello-world

## How to run locally

$ git clone https://github.com/amitkrout/go-hello-world

$ cd go-hello-world

$ docker build -t hello_go_http .

$ docker run -p 8080:8080 -t hello_go_http

Server running (port=8080), route: http://localhost:8080/helloworld

Access the URL from your browser - http://localhost:8080/helloworld. It should display `Hello, World!`