package main

import (
	"fmt"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/cmd/web"
)

func main() {

	// Start the API and Admin servers in the same goroutine
	go func() {
		fmt.Println("Starting API and Admin server on the same server...")
		web.StartServer()
	}()
	// Block the main function to keep the program running
	select {}
}
