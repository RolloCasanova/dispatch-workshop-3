package usecase

import (
	"fmt"
	"log"

	"github.com/RolloCasanova/dispatch-workshop-3/model"
)

// EmployeeService is the interface that wraps the service's methods
// GetAllEmployees, GetEmployeeByID, and CreateEmployee.
type EmployeeService interface {
	GetAllEmployees() (model.Employees, error)
	GetEmployeeByID(id int) (*model.Employee, error)
	CreateEmployee(e model.Employee) error
}

// EmployeeUSecase implements EmployeeService interface.
type EmployeeUsecase struct {
	service EmployeeService
}

// New returns a new EmployeeUsecase instance.
func New(s EmployeeService) *EmployeeUsecase {
	log.Println("In usecase - NewEmployeeUsecase")

	return &EmployeeUsecase{
		service: s,
	}
}

// GetAllEmployees calls the service to returns all employees.
func (eu *EmployeeUsecase) GetAllEmployees() (model.Employees, error) {
	log.Println("In usecase - GetAllEmployees")

	employees, err := eu.service.GetAllEmployees()
	if err != nil {
		return nil, fmt.Errorf("getting all employees from usecase: %v", err)
	}

	return employees, nil
}

// GetEmployeeByID calls the service to returns an employee by its ID.
func (eu *EmployeeUsecase) GetEmployeeByID(id int) (*model.Employee, error) {
	log.Println("In usecase - GetEmployeeByID")

	employee, err := eu.service.GetEmployeeByID(id)
	if err != nil {
		return nil, fmt.Errorf("getting employee by id from usecase: %v", err)
	}

	return employee, nil
}

// CreateEmployee calls the service to create an employee.
func (eu *EmployeeUsecase) CreateEmployee(e model.Employee) error {
	log.Println("In usecase - CreateEmployee")

	err := eu.service.CreateEmployee(e)
	if err != nil {
		return fmt.Errorf("creating employee in usecase: %v", err)
	}

	return nil
}
