package service

import "github.com/dedavidsamya/Pet-Banking/domain"

//The customer service is a primary port.
//in hexagonal architecture a port allows external adapters to connect to domain (core business) logic.
//The customer service is the port that will expose functionality of the core to the external world (user).

// Here we will forward the request which arrives on the primary port, to the secondary port.
// The request arrives on Service (primary port) and we will forward it to Repository (secondary port)
// (which will then retrieve data from the Mock DB (customer repository stub)
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// this func has a receiver, and this means it is acting ON the DefaultCustomerService
// as we can see above, the default customer service contains a repo (which is the customer repository, inside domain).
// So this func acts on the CustomerRepository, and calls the FindAll func which lives inside it.
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// here we create an instance of default customer service and assign the repo (customer repository) to it
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
