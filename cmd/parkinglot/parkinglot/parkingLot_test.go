package parkinglot

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNew_ShouldReturnError_WhenNumberOfSlotsIsNotAPositiveNumber(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(0)

	assert.Nil(parkingLot)
	assert.NotNil(err)
}
  
func TestNew_ShouldReturnParkingLot_WhenNumberOfSlotsIsPositive(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(5)

	assert.Nil(err)
	assert.Equal(5, parkingLot.TotalSlots)
	assert.Equal(5, parkingLot.EmptySlots)
	assert.Equal([]int{1, 2, 3, 4, 5}, parkingLot.SlotsAvailable)
}

func TestPark_ShouldReturnError_WhenSlotsAreNotAvailable(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(1)
	slot, err := parkingLot.Park()

	assert.Nil(err)
	assert.Equal(1, slot)

	slot, err = parkingLot.Park()
	assert.NotNil(err)
}

func TestPark_ShouldReturnASlot_WhenSlotsAreAvailable(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(5)
	slot, err := parkingLot.Park()

	assert.Nil(err)
	assert.Equal(1, slot)

	slot, err = parkingLot.Park()
	assert.Nil(err)
	assert.Equal(2, slot)
}

func TestUnPark_ShouldReturnError_WhenGivenSlotNumberIsNotInRange(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(5)

	err = parkingLot.UnPark(0)
	assert.NotNil(err)
	assert.Equal("Invalid slot number", err.Error())

	err = parkingLot.UnPark(6)
	assert.NotNil(err)
	assert.Equal("Invalid slot number", err.Error())
}

func TestUnPark_ShouldReturnError_WhenGivenSlotNumberIsAlreadyEmpty(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(5)

	err = parkingLot.UnPark(4)
	assert.NotNil(err)
	assert.Equal("Slot already empty", err.Error())
}

func TestUnPark_ShouldMakeSlotNumberAvailable_WhenGivenSlotNumberIsFilled(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(5)
	_, err = parkingLot.Park()
	_, err = parkingLot.Park()
	_, err = parkingLot.Park()
	_, err = parkingLot.Park()

	err = parkingLot.UnPark(2)
	assert.Nil(err)
	assert.Equal(2, parkingLot.EmptySlots)
	assert.Equal(5, parkingLot.TotalSlots)
	assert.Equal([]int{2, 5}, parkingLot.SlotsAvailable)

	err = parkingLot.UnPark(4)
	assert.Nil(err)
	assert.Equal(3, parkingLot.EmptySlots)
	assert.Equal(5, parkingLot.TotalSlots)
	assert.Equal([]int{2, 4, 5}, parkingLot.SlotsAvailable)
}

func TestPark_ShouldReturnNearestSlot_WhenSlotsAreAvailable(t *testing.T) {
	assert := assert.New(t)
	parkingLot, err := New(5)
	
	slotNumber, err := parkingLot.Park()
	assert.Equal(1, slotNumber)
	assert.Nil(err)
	slotNumber, err = parkingLot.Park()
	assert.Equal(2, slotNumber)
	assert.Nil(err)
	slotNumber, err = parkingLot.Park()
	assert.Equal(3, slotNumber)
	assert.Nil(err)
	slotNumber, err = parkingLot.Park()
	assert.Equal(4, slotNumber)
	assert.Nil(err)

	err = parkingLot.UnPark(2)
	err = parkingLot.UnPark(4)

	slotNumber, err = parkingLot.Park()
	assert.Equal(2, slotNumber)
	assert.Nil(err)

	slotNumber, err = parkingLot.Park()
	assert.Equal(4, slotNumber)
	assert.Nil(err)
}