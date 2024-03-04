package main

import (
	"fmt"

	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/routes"
)

func main() {
	route := routes.SetupRouter()
	port := 8080
	address := fmt.Sprintf("localhost:%d", port)
	fmt.Println("Server running on ", address)
	route.Run(address)
}
