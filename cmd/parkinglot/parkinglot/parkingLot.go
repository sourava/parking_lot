package parkinglot

import (
	"errors"
	"parking_lot/cmd/parkinglot/models"
)

type ParkingLot struct {
	SlotsAvailable []int
	TotalSlots int
	EmptySlots int
	SlotToCarMap map[int]*models.Car
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
		SlotToCarMap: make(map[int]*models.Car),
	}, nil
}

func (p *ParkingLot) Park(car *models.Car) (int, error) {
	if p.EmptySlots == 0 {
		return 0, errors.New("Sorry, parking lot is full")
	}
	firstSlot := p.SlotsAvailable[0]
	p.SlotsAvailable = p.SlotsAvailable[1:]
	p.EmptySlots = p.EmptySlots - 1
	p.SlotToCarMap[firstSlot] = car
	return firstSlot, nil
}

func (p *ParkingLot) UnPark(slotNumber int) (*models.Car, error) {
	if slotNumber <= 0 || slotNumber > p.TotalSlots {
		return nil, errors.New("Invalid slot number")
	}

	if p.checkIfSlotEmpty(slotNumber) {
		return nil, errors.New("Slot already empty")
	}

	index := p.getIndexToInsert(slotNumber)

	p.SlotsAvailable = append(p.SlotsAvailable, 0)
	copy(p.SlotsAvailable[index+1:], p.SlotsAvailable[index:])
	p.SlotsAvailable[index] = slotNumber

	p.EmptySlots = p.EmptySlots + 1
	
	unparkedCar := p.SlotToCarMap[slotNumber]
	delete(p.SlotToCarMap, slotNumber);
	
	return unparkedCar, nil
}

func (p *ParkingLot) getIndexToInsert(slotNumber int) int {
	for index, slot := range p.SlotsAvailable {
		if slotNumber < slot {
			return index
		}
	}
	return p.EmptySlots
}

func (p *ParkingLot) checkIfSlotEmpty(slotNumber int) bool {
	for _, slot := range p.SlotsAvailable {
		if slotNumber == slot {
			return true
		}
	}
	return false
}