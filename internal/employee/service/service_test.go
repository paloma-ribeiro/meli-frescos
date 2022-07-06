package service

import (
	"context"
	"errors"
	"testing"

	"github.com/paloma-ribeiro/meli-frescos/internal/employee/domain"
	"github.com/paloma-ribeiro/meli-frescos/internal/employee/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockEmployeeRepository := new(mocks.EmployeeRepository)
	mockEmployee := &domain.Employee{
		ID:           1,
		CardNumberId: "123a",
		FirstName:    "Maria",
		LastName:     "Do Bairro",
		WarehouseId:  12,
	}

	t.Run("success", func(t *testing.T) {
		mockEmployeeRepository.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(mockEmployee, nil).Once()

		service := NewEmployeeService(mockEmployeeRepository)

		employee, err := service.Create(context.TODO(), mockEmployee)

		assert.NoError(t, err)
		assert.Equal(t, "123a", employee.CardNumberId)

		mockEmployeeRepository.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockEmployeeRepository.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(&domain.Employee{}, errors.New("failed to create")).Once()

		service := NewEmployeeService(mockEmployeeRepository)

		_, err := service.Create(context.TODO(), mockEmployee)

		assert.Error(t, err)

		mockEmployeeRepository.AssertExpectations(t)
	})

}
