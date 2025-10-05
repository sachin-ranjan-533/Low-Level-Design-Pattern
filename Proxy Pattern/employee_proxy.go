package main

import "fmt"

type EmployeeProxy struct {
	employee *Employee
	isAdmin  bool
}

func NewEmployeeProxy(isAdmin bool) *EmployeeProxy {
	return &EmployeeProxy{
		employee: &Employee{},
		isAdmin:  isAdmin,
	}
}

func (proxy *EmployeeProxy) InsertEmployee(emp Employee) {
	// Additional logic can be added here (e.g., logging, access control)
	if !proxy.isAdmin {
		fmt.Println("Only admin users can remove employees")
		return
	}
	proxy.employee.InsertEmployee(emp)
}

func (proxy *EmployeeProxy) RemoveEmployee(id int) Employee {
	// Additional logic can be added here (e.g., logging, access control)
	if !proxy.isAdmin {
		fmt.Println("Only admin users can remove employees")
	}
	return proxy.employee.RemoveEmployee(id)
}
