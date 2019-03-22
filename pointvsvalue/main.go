package main

import (
	"fmt"
	"pointvsvalue/pvsv"
)




func main() {
	e := pvsv.Employee{
		Name: "Mark Andrew",
		Age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.Name)
	e.ChangeName("Shrikar vaitala")
	fmt.Printf("\nEmployee name after change: %s", e.Name)

	fmt.Printf("\n\nEmployee age before change: %d", e.Age)
	e.ChangeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.Age)
}
