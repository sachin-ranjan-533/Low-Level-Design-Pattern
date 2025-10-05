package main

type OrdinaryCar struct {
	driveInterface DriveInterface
}

func (oc *OrdinaryCar) setDriveInterface(di DriveInterface) {
	oc.driveInterface = di
}

func (oc *OrdinaryCar) performDrive() {
	oc.driveInterface.drive()
}
