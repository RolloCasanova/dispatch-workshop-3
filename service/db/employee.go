package db

import (
	"errors"

	"github.com/RolloCasanova/dispatch-workshop-3/model"
	"github.com/sirupsen/logrus"
)

type dbService struct {
	logger *logrus.Logger
	data   employeesMap
}

func New(l *logrus.Logger) dbService {
	return dbService{
		logger: l,
		data:   employeeData,
	}
}

type employeesMap map[int]model.Employee

var employeeData = employeesMap{
	1: {
		ID:   1,
		Name: "Roboam",
	},
	2: {
		ID:   2,
		Name: "Miguel",
	},
}

func (dbs dbService) GetAllEmployees() (model.Employees, error) {
	dbs.logger.Info("In db.GetAllEmployees")

	if dbs.data == nil {
		dbs.logger.Error("empty map")

		return nil, errors.New("service: db: empty data")
	}

	employees := make(model.Employees, 0, len(employeeData))

	for _, employee := range employeeData {
		employees = append(employees, employee)
	}

	dbs.logger.WithField("employees", employees).Info("employees retrieved successfully")
	return employees, nil
}
