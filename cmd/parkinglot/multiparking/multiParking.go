package multiparking

import (
	"parking_lot/cmd/parkinglot/parkinglot"
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
