package repository

import (
	"github.com/jmoiron/sqlx"
)

type accountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return accountRepository{db: db}
}

func (r accountRepository) Create(acc Account) (*Account, error) {
	query := "insert into accounts(customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"
	rs, err := r.db.Exec(query,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	)
	if err != nil {
		return nil, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return nil, err
	}
	acc.AccountID = int(id)

	return &acc, nil
}

func (r accountRepository) GetAll(custID int) ([]Account, error) {
	query := "select accound_id, customer_id, opening_date, account_type, amount, status from accounts where customer_id = ?"
	accounts := []Account{}
	err := r.db.Select(&accounts, query, custID)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
