package main

func main() {
	vehicleFactory := &VehicleFactory{}

	car := vehicleFactory.CreateVehicle("car")
	car.drive()

	bike := vehicleFactory.CreateVehicle("bike")
	bike.drive()

	defaultVehicle := vehicleFactory.CreateVehicle("plane")
	defaultVehicle.drive()
}
