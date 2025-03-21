package main

import "github.com/marou9916/expense-tracker-api.git/routes"

func main() {
	// Initialize Gin engine with routes setup
	router := routes.SetupRoutes()

	// Start the server on port 8080
	router.Run(":8080")
}
