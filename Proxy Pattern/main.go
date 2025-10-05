package main

func main() {
	// Client code can use EmployeeProxy to interact with Employee
	proxy := NewEmployeeProxy(true)
	proxy.InsertEmployee(Employee{})
	proxy.RemoveEmployee(1)
}
