package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/asxar/backend-test/internal/test"
	"github.com/asxar/backend-test/internal/user"
)

func main() {
	http.HandleFunc("/", test.IndexHandler)
	http.HandleFunc("/user", user.HandleUserRequest)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
