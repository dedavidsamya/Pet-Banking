package main

import "github.com/dedavidsamya/Pet-Banking/app"

func main() {
	//This func is what starts the server, and handles the requests made to the endpoints. It's located in app.go
	app.Start()
}
