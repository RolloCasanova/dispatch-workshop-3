package db

import (
	"github.com/RolloCasanova/dispatch-workshop-3/model"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

var employeeData = map[int]model.Employee{
	1: {
		ID:   1,
		Name: "Roboam",
	},
	2: {
		ID:   2,
		Name: "Miguel",
	},
}

func GetAllEmployees() ([]model.Employee, error) {
	logger.Info("In db.GetAllEmployees")

	employees := make([]model.Employee, 0, len(employeeData))

	for _, employee := range employeeData {
		employees = append(employees, employee)
	}

	return employees, nil
}
