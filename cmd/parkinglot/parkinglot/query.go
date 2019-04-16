package parkinglot

import (
	"errors"
	"parking_lot/cmd/parkinglot/models"
	"parking_lot/cmd/parkinglot/constants"
)

type Query struct {
	slotToCarMap map[int]*models.Car
	colorToSlotsMap map[string][]int
	registrationNumberToSlotMap map[string]int
}

func NewQuery() *Query {
	return &Query {
		slotToCarMap: make(map[int]*models.Car),
		colorToSlotsMap: make(map[string][]int),
		registrationNumberToSlotMap: make(map[string]int),
	}
}

func (q *Query) Add(car *models.Car, slotNumber int) {
	if value, present := q.colorToSlotsMap[car.Color()]; present {
		q.colorToSlotsMap[car.Color()] = append(value, slotNumber)
	} else {
		q.colorToSlotsMap[car.Color()] = []int{slotNumber}
	}
	q.registrationNumberToSlotMap[car.RegistrationNumber()] = slotNumber
}

func (q *Query) Remove(car *models.Car, slotNumber int) error {
	remainingSlots, err := deleteElementFromSlice(q.colorToSlotsMap[car.Color()], slotNumber)
	
	if err != nil {
		return err
	}
	if _, present := q.registrationNumberToSlotMap[car.RegistrationNumber()]; !present {
		return errors.New(constants.NotFound)
	}
	
	q.colorToSlotsMap[car.Color()] = remainingSlots
	delete(q.registrationNumberToSlotMap, car.RegistrationNumber())
	return nil
}

func (q *Query) SlotNumbersForCarsWithColor(color string) ([]int, error){
	if value, present := q.colorToSlotsMap[color]; present {
		return value, nil
	}
	return nil, errors.New(constants.NotFound)
}

func (q *Query) SlotNumberForRegistrationNumber(registrationNumber string) (int, error) {
	if value, present := q.registrationNumberToSlotMap[registrationNumber]; present {
		return value, nil
	}
	return -1, errors.New(constants.NotFound)
}

func deleteElementFromSlice(slice []int, element int) ([]int, error) {
	position := -1
	for index, value := range(slice) {
		if value == element {
			position = index
			break
		}
	}
	if position != -1 {
		return append(slice[:position], slice[position+1:]...), nil
	}
	return nil, errors.New(constants.NotFound)
}