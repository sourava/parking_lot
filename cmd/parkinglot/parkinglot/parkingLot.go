package parkinglot

import (
	"strconv"
	"sort"
	"errors"
	"parking_lot/cmd/parkinglot/models"
	"parking_lot/cmd/parkinglot/constants"
)

type ParkingLot struct {
	slotsAvailable []int
	totalSlots int
	emptySlots int
	slotToCarMap map[int]*models.Car
	query *Query
}

func New(numberOfSlots int) (*ParkingLot, error) {
	if numberOfSlots <= 0 {
		return nil, errors.New(constants.CreateParkingLotError)
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
		query: NewQuery(),
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
		return 0, errors.New(constants.ParkError)
	}
	firstSlot := p.slotsAvailable[0]
	p.slotsAvailable = p.slotsAvailable[1:]
	p.emptySlots = p.emptySlots - 1
	p.slotToCarMap[firstSlot] = car
	p.query.Add(car, firstSlot)
	return firstSlot, nil
}

func (p *ParkingLot) UnPark(slotNumber int) (*models.Car, error) {
	if slotNumber <= 0 || slotNumber > p.totalSlots {
		return nil, errors.New(constants.UnParkInvalidSlotError)
	}

	if p.checkIfSlotEmpty(slotNumber) {
		return nil, errors.New(constants.UnParkSlotAlreadyEmptyError)
	}

	err := p.query.Remove(p.slotToCarMap[slotNumber], slotNumber)
	if err != nil {
		return nil, err
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

func (p *ParkingLot) SlotNumbersForCarsWithColor(color string) ([]int, error) {
	return p.query.SlotNumbersForCarsWithColor(color)
}

func (p *ParkingLot) SlotNumberForRegistrationNumber(registrationNumber string) (int, error) {
	return p.query.SlotNumberForRegistrationNumber(registrationNumber)
}

func (p *ParkingLot) RegistrationNumbersForCarsWithColor(color string) ([]string, error) {
	registrationNumbers := []string{}
	slotNumbers, err := p.SlotNumbersForCarsWithColor(color)
	if err != nil {
		return nil, err
	}
	for _, value := range(slotNumbers) {
		registrationNumbers = append(registrationNumbers, p.slotToCarMap[value].RegistrationNumber())
	}
	return registrationNumbers, nil
}

func (p *ParkingLot) Status() [][]string {
	keys := []int{}
	details := [][]string{}
    for key := range p.slotToCarMap {
        keys = append(keys, key)
	}
	
	sort.Ints(keys)
	
	for _, key := range keys {
		details = append(details, []string{strconv.Itoa(key), p.slotToCarMap[key].RegistrationNumber(), p.slotToCarMap[key].Color()})
	}
	return details
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