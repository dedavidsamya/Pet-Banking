package domain

// This is the adapter of the secondary port (Customer Repository)
// this is a stub used to populate a dummy repository with dummy data, used to test / mock the functionality of the app
// here, it "mocks" a database. it acts as if it was a real DB but it's just hardcoded data coming from outside the business logic
type CustomerRepositoryStub struct {
	customers []Customer
}

// this func acts on the customer repository stub, so it will return data that the stub gives us
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1234", Name: "Samya", City: "Berlin", DateOfBirth: "19.05.1950", Zipcode: "12345"},
		{Id: "4321", Name: "Rob", City: "Berlin", DateOfBirth: "10.06.1980", Zipcode: "12345"},
		{Id: "4567", Name: "Peter", City: "Berlin", DateOfBirth: "15.09.1982", Zipcode: "12345"},
		{Id: "0987", Name: "Pipa", City: "Berlin", DateOfBirth: "02.08.2023", Zipcode: "12345"},
	}
	return CustomerRepositoryStub{customers}
}
