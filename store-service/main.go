package main

import (
	"fmt"

	"com.ekazeker/store-service/routes"
)

func main() {

	fmt.Println("Starting store-service")

	routes.HandleRequests()

}
