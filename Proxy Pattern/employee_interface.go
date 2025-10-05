package main

type EmployeeInterface interface {
	InsertEmployee(emp Employee)
	RemoveEmployee(id int) Employee
}
