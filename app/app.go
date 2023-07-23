package app

import (
	"github.com/dedavidsamya/Pet-Banking/domain"
	"github.com/dedavidsamya/Pet-Banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//this is the logic that will start the application
//it will define the routes based on defined endpoints,
//it will handle the requests made to those endpoints through a multiplexer,
//and it will start the server

func Start() {

	//this is a Multiplexer, it will handle requests to the endpoints assigning them (matching) to the registered routes
	router := mux.NewRouter()
	//We just need to pass the endpoint "/greet" to the Handlefunc, and the function greet,
	//and the multiplexer will handle the request and direct it to the appropriate route.

	//here is the wiring of the application (connecting primary and secondary ports to the domain core logic)
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//DEFINING ROUTES
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id}", getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	// This is the http package function that will start the server. I don't need to pass a handler because
	//the Handlefunc already took care of it with the default multiplexer. So I just pass the address where
	//the server is supposed to receive requests.

	//STARTING SERVER
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
