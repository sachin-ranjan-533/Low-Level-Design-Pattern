package factory

import "abstract-factory-pattern/vehicle"

type OrdinaryFactory struct{}

func (of *OrdinaryFactory) GetVehicle(vehicleType string) vehicle.Vehicle {
	switch vehicleType {
	case "Car":
		return &vehicle.OrdinaryCar{}
	case "Bike":
		return &vehicle.OrdinaryBike{}
	default:
		return nil
	}
}
