package storage

import "errors"

var (
	ErrBillExists   = errors.New("bill already exists")
	ErrBillNotFound = errors.New("bill not found")
)

const (
	CreateBill = `
		INSERT INTO bills (
		    id,
			bill_name,
			current_sum,
		    bill_type
		) VALUES (
			$1, $2, $3, $4
		) RETURNING id, bill_name, current_sum, bill_type
	`

	CreateAccount = `
		INSERT INTO accounts (
		    bill_id, 
		    bill_status
		) VALUES (
		    $1, $2
		) RETURNING bill_status
	`

	CreateGoals = `
		INSERT INTO goals (
		    bill_id, 
		    goal_sum,
		    date,
		    monthly_payment
		) VALUES (
		    $1, $2, $3, $4
		) RETURNING goal_sum, date, monthly_payment
		`
)

const (
	// HACK: подумать над аккуратностью запроса
	GetOneBill = `
		SELECT id, bill_name, current_sum, bill_type
		FROM bills 
		WHERE id::text = $1
		LIMIT 1;
	`

	GetOneAccount = `
		SELECT bill_status
		FROM accounts
		WHERE bill_id::text = $1
		LIMIT 1 
	`

	GetOneGoal = `
		SELECT goal_sum, date, monthly_payment
		FROM goals
		WHERE bill_id::text = $1
		LIMIT 1
	`

	GetSomeBills = `
		SELECT id, bill_name, current_sum, bill_type
		FROM bills
		WHERE bill_type = $1
		ORDER BY id
	`
	GetAllAccounts = `SELECT bill_status FROM accounts ORDER BY bill_id`
	GetAllGoals    = `
		SELECT goal_sum, date, monthly_payment
		FROM goals
		ORDER BY bill_id
	`
)

const (
	UpdateBill = `
		UPDATE bills
		SET bill_name = $2, current_sum = $3
		WHERE id = $1
		RETURNING bill_name, current_sum, bill_type;
	`
	UpdateAccount = `
		UPDATE accounts
		SET bill_status = $2
		WHERE bill_id = $1
		RETURNING bill_status;
	`
	UpdateGoal = `
		UPDATE goals
		SET goal_sum = $2, date = $3, monthly_payment = $4
		WHERE bill_id = $1
		RETURNING goal_sum, date, monthly_payment;
	`
)

const (
	DeleteBill    = `DELETE FROM bills WHERE id = $1 RETURNING bill_type;`
	DeleteAccount = `DELETE FROM accounts WHERE bill_id = $1;`
	DeleteGoal    = `DELETE FROM goals WHERE bill_id = $1;`
)

const (
	VerifyBill = `SELECT bill_type FROM bills WHERE id::text = $1 LIMIT 1;`
)
