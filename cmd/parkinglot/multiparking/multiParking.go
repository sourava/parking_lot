package multiparking

import (
	"parking_lot/cmd/parkinglot/parkinglot"
	"parking_lot/cmd/parkinglot/models"
)

type MultiParking struct {
	parkingLots []*parkinglot.ParkingLot
	strategy    string
}

func New() *MultiParking {
	return &MultiParking{
		parkingLots: []*parkinglot.ParkingLot{},
		strategy: "",
	}
}

func (m *MultiParking) AddParkingLot(parkingLot *parkinglot.ParkingLot) {
	m.parkingLots = append(m.parkingLots, parkingLot)
}

func (m *MultiParking) AddParkingStrategy(strategy string) {
	m.strategy = strategy
}

func (m *MultiParking) Park(car *models.Car) (int, error) {
	return m.parkingLots[0].Park(car)
}
