package mariadb

import (
	"context"
	"database/sql"

	"github.com/paloma-ribeiro/meli-frescos/internal/employee/domain"
)

type mariadbRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(de *sql.DB) domain.EmployeeRepository {
	return mariadbRepository{db: db}
}

func (m mariadbRepository) GetAll(ctx context.Context) (*[]domain.Employee, error) {
	rows, err := m.db.QueryContext(ctx, "SELECT * FROM employees")

	if err != nil {
		return &[]domain.Employee{}, err
	}

	defer rows.Close()

	var employees []domain.Employee

	for rows.Next() {
		var employee domain.Employee

		if err := rows.Scan(
			&employee.ID,
			&employee.CardNumberId,
			&employee.FirstName,
			&employee.LastName,
			&employee.WarehouseId,
		); err != nil {
			return &[]domain.Employee{}, err
		}

		employees = append(employees, employee)
	}

	return &employees, nil
}

func (m mariadbRepository) GetById(ctx context.Context, id int64) (*domain.Employee, error) {
	row := m.db.QueryRowContext(ctx, "SELECT * FROM employees WHERE ID = ?", id)

	var employee domain.Employee

	if err := row.Scan(
		&employee.ID,
		&employee.CardNumberId,
		&employee.FirstName,
		&employee.LastName,
		&employee.WarehouseId,
	); err != nil {
		return &employee, err
	}

	return &employee, nil
}
