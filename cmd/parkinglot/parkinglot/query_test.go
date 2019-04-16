package parkinglot

import (
	"testing"

	"parking_lot/cmd/parkinglot/models"

	"github.com/stretchr/testify/assert"
)

func TestSlotNumbersForCarsWithColor_ShouldReturnError_WhenNoSlotsFilledWithCarHavingGivenColor(t *testing.T) {
	assert := assert.New(t)

	query := NewQuery()

	slotNumbers, err := query.SlotNumbersForCarsWithColor("White")
	assert.Nil(slotNumbers)
	assert.NotNil(err)
	assert.Equal("Not found", err.Error())
}

func TestSlotNumbersForCarsWithColor_ShouldReturnSlotNumbers_WhenSlotsAreFilledWithCarHavingGivenColor(t *testing.T) {
	assert := assert.New(t)

	query := NewQuery()

	car := models.NewCar("KA-01-HH-1234", "White")
	query.Add(car, 1)

	slotNumbers, err := query.SlotNumbersForCarsWithColor("White")
	assert.Equal([]int{1}, slotNumbers)
	assert.Nil(err)
}

func TestSlotNumberForRegistrationNumber_ShouldReturnError_WhenNoSlotFilledWithCarHavingGivenRegistrationNumber(t *testing.T) {
	assert := assert.New(t)

	query := NewQuery()

	_, err := query.SlotNumberForRegistrationNumber("KA-01-HH-1234")
	assert.NotNil(err)
	assert.Equal("Not found", err.Error())
}

func TestSlotNumberForRegistrationNumber_ShouldReturnSlotNumber_WhenSlotFilledWithCarHavingGivenRegistrationNumber(t *testing.T) {
	assert := assert.New(t)

	query := NewQuery()

	car := models.NewCar("KA-01-HH-1234", "White")
	query.Add(car, 1)

	slotNumber, err := query.SlotNumberForRegistrationNumber("KA-01-HH-1234")
	assert.Equal(1, slotNumber)
	assert.Nil(err)
}