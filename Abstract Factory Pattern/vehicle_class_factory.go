package main

type VehicleClassFactory interface {
	GetVehicle(vehicleType string) Vehicle
}
