package parkinglot

import (
	"errors"
	"parking_lot/cmd/parkinglot/models"
)

type ParkingLot struct {
	slotsAvailable []int
	totalSlots int
	emptySlots int
	slotToCarMap map[int]*models.Car
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
		slotsAvailable: slotsAvailable,
		totalSlots: numberOfSlots,
		emptySlots: numberOfSlots,
		slotToCarMap: make(map[int]*models.Car),
	}, nil
}

func (p *ParkingLot) TotalSlots() int {
	return p.totalSlots
}

func (p *ParkingLot) EmptySlots() int {
	return p.emptySlots
}

func (p *ParkingLot) Park(car *models.Car) (int, error) {
	if p.emptySlots == 0 {
		return 0, errors.New("Sorry, parking lot is full")
	}
	firstSlot := p.slotsAvailable[0]
	p.slotsAvailable = p.slotsAvailable[1:]
	p.emptySlots = p.emptySlots - 1
	p.slotToCarMap[firstSlot] = car
	return firstSlot, nil
}

func (p *ParkingLot) UnPark(slotNumber int) (*models.Car, error) {
	if slotNumber <= 0 || slotNumber > p.totalSlots {
		return nil, errors.New("Invalid slot number")
	}

	if p.checkIfSlotEmpty(slotNumber) {
		return nil, errors.New("Slot already empty")
	}

	index := p.getIndexToInsert(slotNumber)

	p.slotsAvailable = append(p.slotsAvailable, 0)
	copy(p.slotsAvailable[index+1:], p.slotsAvailable[index:])
	p.slotsAvailable[index] = slotNumber

	p.emptySlots = p.emptySlots + 1
	
	unparkedCar := p.slotToCarMap[slotNumber]
	delete(p.slotToCarMap, slotNumber);
	
	return unparkedCar, nil
}

func (p *ParkingLot) getIndexToInsert(slotNumber int) int {
	for index, slot := range p.slotsAvailable {
		if slotNumber < slot {
			return index
		}
	}
	return p.emptySlots
}

func (p *ParkingLot) checkIfSlotEmpty(slotNumber int) bool {
	for _, slot := range p.slotsAvailable {
		if slotNumber == slot {
			return true
		}
	}
	return false
}