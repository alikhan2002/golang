package main

import (
	"Go/manager"
	"fmt"
)

func main() {
	// Create a manager.
	manager := manager.Manager{
		Position: "CEO",
		Salary:   100000,
		Address:  "123 Main Street",
	}

	// Print the position of the manager.
	fmt.Println(manager.GetPosition()) // "CEO"
}
