package main

import (
	"net/http"

	"github.com/Mersock/golang-hexagonal-architecture/handler"
	"github.com/Mersock/golang-hexagonal-architecture/repository"
	"github.com/Mersock/golang-hexagonal-architecture/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:root@tcp(db:3306)/banking?parseTime=true")
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()

	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{CustomerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)
}
