package main

type VehicleFactory struct{}

func (vf *VehicleFactory) CreateVehicle(vehicleType string) Vehicle {
	switch vehicleType {
	case "car":
		return &Car{}
	case "bike":
		return &Bike{}
	default:
		return &DefaultVehicle{}
	}
}
