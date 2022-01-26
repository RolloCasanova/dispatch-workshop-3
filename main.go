package main

import (
	"log"
	"net/http"

	"github.com/RolloCasanova/dispatch-workshop-3/router"
)

func main() {
	r := router.Setup()

	log.Println("Server is running on port 8080...")
	http.ListenAndServe("localhost:8080", r)
}
