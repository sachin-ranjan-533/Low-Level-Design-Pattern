package main

type VehicleFactory struct {
	vehicleClassFactory VehicleClassFactory
}

func (vf *VehicleFactory) GetVehicle(class string, vehicleType string) Vehicle {
	if class == "LUXURY" {
		vf.vehicleClassFactory = &VehicleLuxuryFactory{}
	} else if class == "ORDINARY" {
		vf.vehicleClassFactory = &VehicleOrdinaryFactory{}
	}
	return vf.vehicleClassFactory.GetVehicle(vehicleType)
}
