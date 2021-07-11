package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{CustumerID: 1001, Name: "Leonel", City: "Barcelona", ZipCode: "1001", Status: 1},
		{CustumerID: 1001, Name: "Neymar", City: "Paris", ZipCode: "3001", Status: 1},
	}
	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {

	for _, customer := range r.customers {
		if customer.CustumerID == id {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}
