package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte("Hello World!"))
	}).Methods(http.MethodGet)

	log.Println("Server is running on port 8080...")
	http.ListenAndServe("localhost:8080", router)

}
