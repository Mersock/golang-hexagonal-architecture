package repository

type Custumer struct {
	CustumerID  int    `db:"customer_id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	Status      string `db:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Custumer, error)
	GetById(int) (*Custumer, error)
}
