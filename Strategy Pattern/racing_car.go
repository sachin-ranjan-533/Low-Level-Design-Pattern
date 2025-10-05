package main

type RacingCar struct {
	driveInterface DriveInterface
}

func (rc *RacingCar) setDriveInterface(di DriveInterface) {
	rc.driveInterface = di
}

func (rc *RacingCar) performDrive() {
	rc.driveInterface.drive()
}
