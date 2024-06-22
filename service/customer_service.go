package service

import (
	"database/sql"

	"natthan.com/go-play/errs"
	"natthan.com/go-play/logs"
	"natthan.com/go-play/repository"
)

// Business - Adapter

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	// Don't Return this!
	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	// Return response instead
	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}
	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {
		// Business Level Error
		if err == sql.ErrNoRows {
			return nil, errs.NewNotfoundError("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custReponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custReponse, nil
}
