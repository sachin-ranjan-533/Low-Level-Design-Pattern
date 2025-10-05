package main

func main() {
	vf := &VehicleFactory{}
	vf.GetVehicle("LUXURY", "Car").drive()
	vf.GetVehicle("LUXURY", "Bike").drive()
	vf.GetVehicle("ORDINARY", "Car").drive()
	vf.GetVehicle("ORDINARY", "Bike").drive()
}
