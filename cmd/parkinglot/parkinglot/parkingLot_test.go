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