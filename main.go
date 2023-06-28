package main

import (
	"fmt"

	setup "localhost/database"

	"localhost/middleware"
)

func main() {

	fmt.Print("Hello World")
	err := setup.Connect()
	if err != nil {
		panic(err)
	}
	middleware.Routers()
}
