package main

func main() {
	abstractFactory := &AbstractFactory{}

	luxuryFactory := abstractFactory.CreateFactory("Luxury")
	luxuryCar := luxuryFactory.GetVehicle("Car")
	luxuryCar.Drive()
	luxuryBike := luxuryFactory.GetVehicle("Bike")
	luxuryBike.Drive()

	ordinaryFactory := abstractFactory.CreateFactory("Ordinary")
	ordinaryCar := ordinaryFactory.GetVehicle("Car")
	ordinaryCar.Drive()
	ordinaryBike := ordinaryFactory.GetVehicle("Bike")
	ordinaryBike.Drive()
}
