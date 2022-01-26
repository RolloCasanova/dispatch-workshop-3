package usecase

import (
	"fmt"

	"github.com/RolloCasanova/dispatch-workshop-3/model"

	"github.com/sirupsen/logrus"
)

type dbService interface {
	GetAllEmployees() (model.Employees, error)
}
type employeeUseCase struct {
	logger    *logrus.Logger
	dbService dbService
}

func New(l *logrus.Logger, ds dbService) employeeUseCase {
	return employeeUseCase{
		logger:    l,
		dbService: ds,
	}
}

func (eu employeeUseCase) GetAllEmployees() (model.Employees, error) {
	eu.logger.Info("In usecase.GetAllEmployees")

	employees, err := eu.dbService.GetAllEmployees()
	if err != nil {
		eu.logger.Error("something bad happened in usecase.GetAllEmployees")

		return nil, fmt.Errorf("usecase: getting all employees: %w", err)
	}

	return employees, nil
}
