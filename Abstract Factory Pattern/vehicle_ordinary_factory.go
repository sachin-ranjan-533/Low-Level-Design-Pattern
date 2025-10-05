package main

type VehicleOrdinaryFactory struct{}

func (vlf *VehicleOrdinaryFactory) GetVehicle(vehicleType string) Vehicle {
	if vehicleType == "Car" {
		return &OrdinaryCar{}
	} else if vehicleType == "Bike" {
		return &OrdinaryBike{}
	}
	return nil
}
