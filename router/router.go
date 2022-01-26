package router

import (
	"net/http"

	"github.com/RolloCasanova/dispatch-workshop-3/controller"

	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/employees", controller.GetAllEmployees).
		Methods(http.MethodGet)

	return router
}
