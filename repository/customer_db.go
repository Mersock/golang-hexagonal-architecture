package repository

import "github.com/jmoiron/sqlx"

type customerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepository {
	return customerRepository{db: db}
}

func (r customerRepository) GetAll() ([]Custumer, error) {

	return nil, nil
}

func (r customerRepository) GetById(Id int) (*Custumer, error) {
	return nil, nil
}
