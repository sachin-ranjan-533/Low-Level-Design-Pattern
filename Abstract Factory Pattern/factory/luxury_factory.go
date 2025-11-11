package factory

import "abstract-factory-pattern/vehicle"

type LuxuryFactory struct{}

func (lf *LuxuryFactory) GetVehicle(vehicleType string) vehicle.Vehicle {
	switch vehicleType {
	case "Car":
		return &vehicle.LuxuryCar{}
	case "Bike":
		return &vehicle.LuxuryBike{}
	default:
		return nil
	}
}
