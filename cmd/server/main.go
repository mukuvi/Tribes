package main

import (
	"fmt"
	"net/http"

	"github.com/mukuvi/Tribes/internal/routes"
)

func main() {
	r := routes.SetupRoutes()

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
