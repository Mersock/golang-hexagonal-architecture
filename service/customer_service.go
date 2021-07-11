package service

import (
	"database/sql"
	"errors"

	"github.com/Mersock/golang-hexagonal-architecture/logs"
	"github.com/Mersock/golang-hexagonal-architecture/repository"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	result := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustumerID: customer.CustumerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		result = append(result, custResponse)
	}

	return result, nil
}

func (s customerService) GetCustomer(Id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}
		logs.Error(err)
		return nil, err
	}

	result := CustomerResponse{
		CustumerID: customer.CustumerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &result, nil
}
