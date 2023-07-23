package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/dedavidsamya/Pet-Banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

// This func receives a writer and the request. The writer is going to take care of
// writing to the response.
func greet(w http.ResponseWriter, r *http.Request) {
	//here we are writing to the response, by passing the writer and the text which will be written
	fmt.Fprint(w, "Hello World")
}

// func that retrieves all customers
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//here I am creating some instances of Customer, using the type Customer I created above (hardcoded just to test)
	//customers := []Customer{
	//	{Name: "Samya", City: "Berlin", Zipcode: "12345"},
	//	{Name: "Rob", City: "Berlin", Zipcode: "11111"},
	//}

	//here, the customers will come from the GetAllCustomer
	customers, _ := ch.service.GetAllCustomer()

	// When the person making the request wants to specify which format they want the response in,
	//they need to pass the header content-type to the request (doing that on postman for example by adding
	//the header on the request headers field).
	//So I'll add a check on my server to see if who made the request specified this or not, and I'll return the data
	//in the format they requested (here for example xml, and the default would be json)
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		//Here I specify I want the content of the response to be in format Json, so I need to pass the Header content-type
		w.Header().Add("content-type", "application/json")
		//and here I am encoding my customer list to JSON format
		json.NewEncoder(w).Encode(customers)
	}
}

// This func will retrieve one specific entry on the db or memory, based on a request parameter
// the endpoint allows a parameter to be passed, there is a placeholder on the endpoint
// customers/{customer_id}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	//to return a value from a URL segment (or route parameters,
	//which is what replaces this placeholder on the endpoint
	///customers/{customer_id}
	vars := mux.Vars(r)
	//here I am printing the URL segment value (request parameter) on the response
	fmt.Fprint(w, vars["customer_id"])
}

// This func is our method POST func, the request which creates a new entry on our db or memory
func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "post request received")
}
