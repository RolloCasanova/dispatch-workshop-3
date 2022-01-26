package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type controller interface {
	GetAllEmployees(w http.ResponseWriter, r *http.Request)
}

func Setup(c controller) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/employees", c.GetAllEmployees).
		Methods(http.MethodGet)

	return router
}
