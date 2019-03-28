package parkinglot

import (
	"errors"
)

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

func (p *ParkingLot) Park() (int, error) {
	if p.EmptySlots == 0 {
		return 0, errors.New("Sorry, parking lot is full")
	}
	firstSlot := p.SlotsAvailable[0]
	p.SlotsAvailable = p.SlotsAvailable[1:]
	p.EmptySlots = p.EmptySlots - 1
	return firstSlot, nil
}