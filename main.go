package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"natthan.com/go-play/handler"
	"natthan.com/go-play/logs"
	"natthan.com/go-play/repository"
	"natthan.com/go-play/service"
)

func main() {
	initTimezone()
	initConfig()
	db := initDatabase()
	// Create Repo
	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()

	_ = customerRepositoryDB
	_ = customerRepositoryMock

	accountRepositoryDB := repository.NewAccountRepositoryDB(db)
	accountService := service.NewAccountService(accountRepositoryDB)
	accountHandler := handler.NewAccountHandler(accountService)

	customerService := service.NewCustomerService(customerRepositoryMock)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	// customer
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)
	// account
	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.GetAccounts).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.NewAccount).Methods(http.MethodPost)

	log.Printf("Banking service started at port %v", viper.GetInt("app.port"))
	logs.Info("Banking service started at port " + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
}

func initConfig() {
	// Can read config and environment
	// Priority Environment
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	// APP_PORT -> app.port
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimezone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	// Can specify ?parseTime=true to get Time type
	// Datasource Name
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
