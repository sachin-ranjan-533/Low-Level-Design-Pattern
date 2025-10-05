package main

import "fmt"

type Employee struct{}

func (e *Employee) InsertEmployee(emp Employee) {
	// Logic to insert employee into the database
	fmt.Println("Employee inserted")
}

func (e *Employee) RemoveEmployee(id int) Employee {
	// Logic to remove employee from the database
	fmt.Println("Employee removed")
	return Employee{}
}
