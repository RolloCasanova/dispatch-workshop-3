package usecase

import (
	"fmt"

	"github.com/RolloCasanova/dispatch-workshop-3/model"
	"github.com/RolloCasanova/dispatch-workshop-3/service/db"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func GetAllEmployees() (model.Employees, error) {
	logger.Info("In usecase.GetAllEmployees")

	employees, err := db.GetAllEmployees()
	if err != nil {
		logger.Error("something bad happened in usecase.GetAllEmployees")

		return nil, fmt.Errorf("usecase: getting all employees: %w", err)
	}

	return employees, nil
}
