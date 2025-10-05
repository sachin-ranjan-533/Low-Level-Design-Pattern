package main

func main() {
	ordinaryCar := &OrdinaryCar{}
	ordinaryCar.setDriveInterface(&NormalDrive{})
	ordinaryCar.performDrive()

	racingCar := &RacingCar{}
	racingCar.setDriveInterface(&RacingDrive{})
	racingCar.performDrive()

	superCar := &SuperCar{}
	superCar.setDriveInterface(&RacingDrive{})
	superCar.performDrive()
}
