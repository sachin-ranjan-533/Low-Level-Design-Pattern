package main

func main() {
	vehicleFactory := &VehicleFactory{}
	
	car := vehicleFactory.CreateVehicle("car")
	car.drive()
}