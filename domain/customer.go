package domain

// domain object (core business object)
type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// CustomerRepository is a secondary port. The secondary port will allow adapters (e.g. DB or mock) to connect to the business logic and transmit data to it.
// This is a concept of hexagonal architecture and will keep the core business logic adapter-agnostic,
// meaning we can change the DB for another, and the business logic does not need to be touched.
// This CustomerRepository interface implements the FindAll func which is a func that acts on the Stub Repo
// so when we call FindAll, we are calling it on the stub (see customer_repository_stub.go)
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
