package main

type VehicleLuxuryFactory struct{}

func (vlf *VehicleLuxuryFactory) GetVehicle(vehicleType string) Vehicle {
	if vehicleType == "Car" {
		return &LuxuryCar{}
	} else if vehicleType == "Bike" {
		return &LuxuryBike{}
	}
	return nil
}
