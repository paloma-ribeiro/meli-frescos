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
