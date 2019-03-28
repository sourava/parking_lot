package parkinglot

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnError_WhenNumberOfSlotsIsNotAPositiveNumber(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(0)

	assert.Nil(parkingLot)
	assert.NotNil(err)
}
  
func Test_ShouldReturnParkingLotWithGivenNumberOfSlots_WhenNumberOfSlotsIsPositive(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(5)

	assert.Nil(err)
	assert.Equal(5, parkingLot.TotalSlots)
	assert.Equal(5, parkingLot.EmptySlots)
	assert.Equal([]int{1, 2, 3, 4, 5}, parkingLot.SlotsAvailable)
}