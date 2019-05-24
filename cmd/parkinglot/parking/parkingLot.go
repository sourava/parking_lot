package parking

import (
	"parking_lot/cmd/parkinglot/models"
)

type ParkingLot interface {
	Park(car *models.Car) (int, error)
	UnPark(slotNumber int) (*models.Car, error)
}
