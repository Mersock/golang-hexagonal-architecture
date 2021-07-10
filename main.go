package main

import (
	"github.com/Mersock/golang-hexagonal-architecture/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:P@ssw0rd@tcp(mysql-db:3306)/banking")
	if err != nil {
		panic(err)
	}
	customerRepository := repository.NewCustomerRepositoryDB(db)
	_ = customerRepository
}
