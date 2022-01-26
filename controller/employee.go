package controller

import (
	"fmt"
	"net/http"

	"github.com/RolloCasanova/dispatch-workshop-3/model"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type usecase interface {
	GetAllEmployees() (model.Employees, error)
}

type employeeController struct {
	logger  *logrus.Logger
	render  *render.Render
	usecase usecase
}

func New(l *logrus.Logger, r *render.Render, u usecase) *employeeController {
	return &employeeController{
		logger:  l,
		render:  r,
		usecase: u,
	}
}

func (ec employeeController) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	ec.logger.Info("In controller.GetAllEmployees")

	employees, err := ec.usecase.GetAllEmployees()
	if err != nil {
		ec.logger.Error("something failed calling controller.GetAllEmployees")

		err = fmt.Errorf("controller: getting all employees: %w", err)
		ec.render.Text(w, http.StatusInternalServerError, err.Error())

		return
	}

	ec.render.JSON(w, http.StatusOK, employees)
}
