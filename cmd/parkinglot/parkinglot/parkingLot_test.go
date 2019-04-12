package parkinglot

import (
	"testing"

	"parking_lot/cmd/parkinglot/models"

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
	assert.Equal(5, parkingLot.TotalSlots())
	assert.Equal(5, parkingLot.EmptySlots())
}

func TestPark_ShouldReturnError_WhenGivenACarAndSlotsAreNotAvailable(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(1)
	car1 := models.NewCar("KA-01-HH-1234", "White")
	slot, err := parkingLot.Park(car1)

	assert.Nil(err)
	assert.Equal(1, slot)

	car2 := models.NewCar("KA-02-HH-1234", "White")
	slot, err = parkingLot.Park(car2)
	assert.NotNil(err)
}

func TestPark_ShouldReturnASlot_WhenGivenACarAndSlotsAreAvailable(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(5)
	car1 := models.NewCar("KA-01-HH-1234", "White")
	slot, err := parkingLot.Park(car1)

	assert.Nil(err)
	assert.Equal(1, slot)

	car2 := models.NewCar("KA-02-HH-1234", "White")
	slot, err = parkingLot.Park(car2)
	assert.Nil(err)
	assert.Equal(2, slot)
}

func TestUnPark_ShouldReturnError_WhenGivenSlotNumberIsNotInRange(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(5)

	car, err := parkingLot.UnPark(0)
	assert.Nil(car)
	assert.NotNil(err)
	assert.Equal("Invalid slot number", err.Error())

	car, err = parkingLot.UnPark(6)
	assert.Nil(car)
	assert.NotNil(err)
	assert.Equal("Invalid slot number", err.Error())
}

func TestUnPark_ShouldReturnError_WhenGivenSlotNumberIsAlreadyEmpty(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(5)

	car, err := parkingLot.UnPark(4)
	assert.Nil(car)
	assert.NotNil(err)
	assert.Equal("Slot already empty", err.Error())
}

func TestUnPark_ShouldReturnParkedCar_WhenGivenSlotNumberIsFilled(t *testing.T) {
	assert := assert.New(t)

	parkingLot, err := New(5)
	car1 := models.NewCar("KA-01-HH-1234", "White")
	car2 := models.NewCar("KA-02-HH-1234", "White")
	car3 := models.NewCar("KA-03-HH-1234", "White")
	car4 := models.NewCar("KA-04-HH-1234", "White")
	_, err = parkingLot.Park(car1)
	_, err = parkingLot.Park(car2)
	_, err = parkingLot.Park(car3)
	_, err = parkingLot.Park(car4)

	unparkedCar, err := parkingLot.UnPark(2)
	assert.Equal(car2, unparkedCar)
	assert.Nil(err)
	assert.Equal(2, parkingLot.EmptySlots())
	assert.Equal(5, parkingLot.TotalSlots())

	unparkedCar, err = parkingLot.UnPark(4)
	assert.Equal(car4, unparkedCar)
	assert.Nil(err)
	assert.Equal(3, parkingLot.EmptySlots())
	assert.Equal(5, parkingLot.TotalSlots())
}

func TestPark_ShouldReturnNearestSlot_WhenSlotsAreAvailable(t *testing.T) {
	assert := assert.New(t)
	parkingLot, err := New(5)
	car1 := models.NewCar("KA-01-HH-1234", "White")
	car2 := models.NewCar("KA-02-HH-1234", "White")
	car3 := models.NewCar("KA-03-HH-1234", "White")
	car4 := models.NewCar("KA-04-HH-1234", "White")
	
	slotNumber, err := parkingLot.Park(car1)
	assert.Equal(1, slotNumber)
	assert.Nil(err)
	slotNumber, err = parkingLot.Park(car2)
	assert.Equal(2, slotNumber)
	assert.Nil(err)
	slotNumber, err = parkingLot.Park(car3)
	assert.Equal(3, slotNumber)
	assert.Nil(err)
	slotNumber, err = parkingLot.Park(car4)
	assert.Equal(4, slotNumber)
	assert.Nil(err)

	_, err = parkingLot.UnPark(2)
	_, err = parkingLot.UnPark(4)

	slotNumber, err = parkingLot.Park(car2)
	assert.Equal(2, slotNumber)
	assert.Nil(err)

	slotNumber, err = parkingLot.Park(car4)
	assert.Equal(4, slotNumber)
	assert.Nil(err)
}