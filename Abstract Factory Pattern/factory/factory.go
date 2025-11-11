package factory

import "abstract-factory-pattern/vehicle"

type Factory interface {
	GetVehicle(vehicleType string) vehicle.Vehicle
}
