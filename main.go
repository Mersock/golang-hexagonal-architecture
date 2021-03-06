package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Mersock/golang-hexagonal-architecture/handler"
	"github.com/Mersock/golang-hexagonal-architecture/logs"
	"github.com/Mersock/golang-hexagonal-architecture/repository"
	"github.com/Mersock/golang-hexagonal-architecture/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	config "github.com/spf13/viper"
)

var db *sqlx.DB

func init() {
	initConfig()
	initTimeZone()
	db = initDB()
}

func main() {
	router := mux.NewRouter()

	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)
	accountHandler := handler.NewAccountHandler(accountService)

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{CustomerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers/{CustomerID:[0-9]+}/accounts", accountHandler.GetAccounts).Methods(http.MethodGet)
	router.HandleFunc("/customers/{CustomerID:[0-9]+}/accounts", accountHandler.NewAccount).Methods(http.MethodPost)

	logs.Info("Start service at port " + config.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", config.GetInt("app.port")), router)
}

func initConfig() {
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")

	config.AutomaticEnv()
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDB() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.GetString("db.username"),
		config.GetString("db.password"),
		config.GetString("db.host"),
		config.GetInt("db.port"),
		config.GetString("db.database"),
	)

	db, err := sqlx.Open(config.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
