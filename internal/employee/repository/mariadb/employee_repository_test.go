package mariadb

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/paloma-ribeiro/meli-frescos/internal/employee/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockEmployee := &domain.Employee{
		CardNumberId: "123a",
		FirstName:    "Maria",
		LastName:     "Do Bairro",
		WarehouseId:  12,
	}

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(queryCreate)).WithArgs(
			mockEmployee.CardNumberId,
			mockEmployee.FirstName,
			mockEmployee.LastName,
			mockEmployee.WarehouseId,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		employeeRepository := NewMariaDBRepository(db)
		employee, err := employeeRepository.Create(context.TODO(), mockEmployee)

		assert.NoError(t, err)

		assert.Equal(t, "123a", employee.CardNumberId)
	})

	t.Run("failed to create", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(queryCreate)).WithArgs("123a", "Maria", "Do Bairro", 12)

		employeeRepository := NewMariaDBRepository(db)
		_, err = employeeRepository.Create(context.TODO(), mockEmployee)

		assert.Error(t, err)
	})
}
