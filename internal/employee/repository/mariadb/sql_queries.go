package mariadb

const (
	queryGetAll = `SELECT 
	id
	card_number_id
	first_name
	last_name
	warehouse_id
	FROM employees`

	queryGetById = `SELECT 
	id
	card_number_id
	first_name
	last_name
	warehouse_id
	FROM employees
	WHERE ID = ?`

	queryCreate = `INSERT INTO employees 
	(card_number_id, first_name, last_name, warehouse_id) 
	VALUES (?, ?, ?, ?)`
)
