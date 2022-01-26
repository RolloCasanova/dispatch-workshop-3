package main

import (
	"log"
	"net/http"

	"github.com/RolloCasanova/dispatch-workshop-3/controller"
	"github.com/RolloCasanova/dispatch-workshop-3/router"
	"github.com/RolloCasanova/dispatch-workshop-3/service/db"
	"github.com/RolloCasanova/dispatch-workshop-3/usecase"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

func main() {
	l := logrus.New()

	ds := db.New(l)
	u := usecase.New(l, ds)
	c := controller.New(l, render.New(), u)
	r := router.Setup(c)

	log.Println("Server is running on port 8080...")
	http.ListenAndServe("localhost:8080", r)
}
