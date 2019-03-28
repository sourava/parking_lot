package parkinglot

import "errors"

type ParkingLot struct {
	SlotsAvailable []int
	TotalSlots int
	EmptySlots int
}

func New(numberOfSlots int) (*ParkingLot, error) {
	if numberOfSlots <= 0 {
		return nil, errors.New("Enter positive number in numberOfSlots")
	}

	index := 0
	slotsAvailable := make([]int, numberOfSlots)
	for index = 0; index < numberOfSlots; index++ {
		slotsAvailable[index] = index + 1
	}
	return &ParkingLot {
		SlotsAvailable: slotsAvailable,
		TotalSlots: numberOfSlots,
		EmptySlots: numberOfSlots,
	}, nil
}