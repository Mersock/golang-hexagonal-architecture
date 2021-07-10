package repository

import "github.com/jmoiron/sqlx"

type customerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepository {
	return customerRepository{db: db}
}
