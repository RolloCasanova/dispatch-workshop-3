package controller

import (
	"net/http"

	"github.com/RolloCasanova/dispatch-workshop-3/model"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

var (
	logger = logrus.New()
	ren    = render.New()

	employees = model.Employees{
		{ID: 1, Name: "Roboam"},
		{ID: 2, Name: "Miguel"},
	}
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	logger.Info("In controller.GetAllEmployees")

	ren.JSON(w, http.StatusOK, employees)
}
