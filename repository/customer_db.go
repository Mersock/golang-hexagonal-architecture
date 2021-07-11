package repository

import (
	"github.com/jmoiron/sqlx"
)

type customerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) customerRepository {
	return customerRepository{db: db}
}

func (r customerRepository) GetAll() ([]Custumer, error) {
	customers := []Custumer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepository) GetById(Id int) (*Custumer, error) {
	customer := Custumer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id = ?"
	err := r.db.Get(&customer, query, Id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
