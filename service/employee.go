package service

import (
	"log"

	errz "github.com/RolloCasanova/dispatch-workshop-3/errors"
	"github.com/RolloCasanova/dispatch-workshop-3/model"
	"github.com/sirupsen/logrus"
)

// EmployeeMap is an alias for a map of employees.
type EmployeeMap map[int]model.Employee

// employeesOrder is an auxiliary function that helps to sort employeeData by their ID
// as we are using a map we can't ensure the order of the keys is preserved.
var employeesOrder []int = []int{1, 2}

// db is a sample data to be used in the service as a placeholder for the real data.
var db EmployeeMap = map[int]model.Employee{
	1: {
		ID:      1,
		Name:    "1",
		Email:   "1",
		Phone:   "1",
		Address: "1",
	},
	2: {
		ID:      2,
		Name:    "2",
		Email:   "2",
		Phone:   "2",
		Address: "2",
	},
}

// EmployeeServicestruct implements EmployeeService interface.
type EmployeeService struct {
	logger *logrus.Logger
	data   EmployeeMap
}

// New returns a new EmployeeService instance.
func New(em EmployeeMap) *EmployeeService {
	if em == nil {
		em = db
	}

	return &EmployeeService{
		data: em,
	}
}

// GetAllEmployees returns all employees data.
func (es *EmployeeService) GetAllEmployees() (model.Employees, error) {
	log.Println("In service - GetAllEmployees")

	if err := es.dataValidation(); err != nil {
		return nil, err
	}

	// convert data from map to an slice of Employees
	employees := make(model.Employees, 0, len(es.data))

	// preserve the order
	for _, id := range employeesOrder {
		employees = append(employees, es.data[id])
	}

	return employees, nil
}

// GetEmployeeByID returns an employee by its ID.
func (es *EmployeeService) GetEmployeeByID(id int) (*model.Employee, error) {
	log.Println("In service - GetEmployeeByID")

	if err := es.dataValidation(); err != nil {
		return nil, err
	}

	// find the employee in the data
	employee, ok := es.data[id]
	if !ok {
		return nil, errz.ErrNotFound
	}

	return &employee, nil
}

// CreateEmployee creates a new employee and add it to the db
// if it already exists it won't be overwritten.
func (es *EmployeeService) CreateEmployee(e model.Employee) error {
	log.Println("In service - CreateEmployee")

	if err := es.dataValidation(); err != nil {
		return err
	}

	// special handling if employee already exists
	if _, ok := es.data[e.ID]; ok {
		return errz.ErrEmployeeAlreadyExists
	}

	// add employee
	es.data[e.ID] = e

	// update employees order
	employeesOrder = append(employeesOrder, e.ID)

	return nil
}

// dataValidation is an auxiliary function that checks if the data has beem initialized or if it is empty
// returns a matching ServiceError if any of these conditions are met.
func (es *EmployeeService) dataValidation() error {
	log.Println("In service - dataValidation")

	// special handling if data is nil
	if es.data == nil {
		return errz.ErrDataNotInitialized
	}

	// special handling if data is empty
	if len(es.data) == 0 {
		return errz.ErrEmptyData
	}

	return nil
}
