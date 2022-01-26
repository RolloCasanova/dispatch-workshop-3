package controller

import (
	"fmt"
	"net/http"

	"github.com/RolloCasanova/dispatch-workshop-3/usecase"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

var (
	logger = logrus.New()
	ren    = render.New()
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	logger.Info("In controller.GetAllEmployees")

	employees, err := usecase.GetAllEmployees()
	if err != nil {
		logger.Error("something failed calling controller.GetAllEmployees")

		err = fmt.Errorf("controller: getting all employees: %w", err)
		ren.Text(w, http.StatusInternalServerError, err.Error())

		return
	}

	ren.JSON(w, http.StatusOK, employees)
}
