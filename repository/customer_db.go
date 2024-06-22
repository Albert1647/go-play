package repository

import "github.com/jmoiron/sqlx"

// Adapter
// adapter is private

type customerRepositoryDB struct {
	db *sqlx.DB
}

// constructor in OOP
// instantiate only to control initialization
func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepository {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "SELECT cusomer_id, name, date_of_birth, zipcode, status FROM customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "SELECT cusomer_id, name, date_of_birth, zipcode, status FROM customers WHERE customer_id = ?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
