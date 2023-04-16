package main

import (
	"fmt"

	"localhost/setup"

	"localhost/middleware"
)

func main() {

	fmt.Print("Hello World")
	setup.Connect()
	middleware.Routers()
}
