package service

import (
	"context"

	"github.com/paloma-ribeiro/meli-frescos/internal/employee/domain"
)

type employeeService struct {
	repository domain.EmployeeRepository
}

func NewEmployee(er domain.EmployeeRepository) domain.EmployeeService {
	return &employeeService{repository: er}
}

func (e employeeService) GetAll(ctx context.Context) (*[]domain.Employee, error) {
	employees, err := e.repository.GetAll(ctx)

	if err != nil {
		return employees, err
	}

	return employees, nil
}

func (e employeeService) GetById(ctx context.Context, id int64) (*domain.Employee, error) {
	employee, err := e.repository.GetById(ctx, id)

	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (e employeeService) Create(ctx context.Context, employee *domain.Employee) (*domain.Employee, error) {
	employee, err := e.repository.Create(ctx, employee)

	if err != nil {
		return employee, err
	}

	return employee, nil
}