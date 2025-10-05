package main

type SuperCar struct {
	driveInterface DriveInterface
}

func (sc *SuperCar) setDriveInterface(di DriveInterface) {
	sc.driveInterface = di
}

func (sc *SuperCar) performDrive() {
	sc.driveInterface.drive()
}
