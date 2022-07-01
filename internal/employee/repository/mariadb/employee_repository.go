package mariadb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/paloma-ribeiro/meli-frescos/internal/employee/domain"
)

type mariadbRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(de *sql.DB) domain.EmployeeRepository {
	return mariadbRepository{db: db}
}

func (m mariadbRepository) GetAll(ctx context.Context) (*[]domain.Employee, error) {
	rows, err := m.db.QueryContext(ctx, queryGetAll)

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

	row := m.db.QueryRowContext(ctx, queryGetById, id)

	var employee domain.Employee

	err := row.Scan(
		&employee.ID,
		&employee.CardNumberId,
		&employee.FirstName,
		&employee.LastName,
		&employee.WarehouseId,
	)

	// ID not found
	if errors.Is(err, sql.ErrNoRows) {
		return &employee, domain.ErrIdNotFound
	}

	// Other errors
	if err != nil {
		return &employee, err
	}

	return &employee, nil
}

func (m mariadbRepository) Create(ctx context.Context, employee *domain.Employee) (*domain.Employee, error) {
	var newEmployee domain.Employee

	result, err := m.db.ExecContext(ctx, queryCreate,
		&newEmployee.CardNumberId,
		&newEmployee.FirstName,
		&newEmployee.LastName,
		&newEmployee.WarehouseId,
	)

	if err != nil {
		return &newEmployee, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return &newEmployee, err
	}

	newEmployee.ID = lastId

	return &newEmployee, nil
}
