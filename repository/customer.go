package repository

// Interface

type Customer struct {
	CustomerID  int    `db:"customer_id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"` //time.Time -> ?parseTime=true
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	Status      int    `db:"status"`
}

// What hold data
type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error) // cannot return nil when use struct, use pointer instead
}
