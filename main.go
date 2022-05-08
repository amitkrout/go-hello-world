package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amitkrout/go-hello-world/pkg"
)

func main() {
	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! - Addition of 5 + 5 is "+fmt.Sprint(pkg.Addition(5, 5))+"\n")
	})
	fmt.Printf("Server running on port=8080, URL: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
