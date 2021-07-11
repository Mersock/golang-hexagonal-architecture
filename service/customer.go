package service

type CustomerResponse struct {
	CustumerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}
