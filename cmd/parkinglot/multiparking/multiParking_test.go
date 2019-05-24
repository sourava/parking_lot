package multiparking

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"parking_lot/cmd/parkinglot/parkinglot"
	"parking_lot/cmd/parkinglot/models"
	"parking_lot/cmd/parkinglot/constants"
)

func TestNew_ShouldReturnAMultiParking(t *testing.T) {
	multiParking := New()

	assert.Equal(t, &MultiParking{
		parkingLots: []*parkinglot.ParkingLot{},
		strategy: "",
	} , multiParking)
}

func TestAddParkingLot_ShouldAddParkingLotToParkingLotList(t *testing.T) {
	multiParking := New()
	parkingLot, _ := parkinglot.New(10)

	multiParking.AddParkingLot(parkingLot)
	assert.Equal(t, &MultiParking{
		parkingLots: []*parkinglot.ParkingLot{
			parkingLot,
		},
		strategy: "",
	}, multiParking)
}

func TestAddParkingStrategy_ShouldAddParkingStrategyToMultiParking(t *testing.T) {
	multiParking := New()
	multiParking.AddParkingStrategy(constants.EvenDistribution)

	assert.Equal(t, &MultiParking{
		parkingLots: []*parkinglot.ParkingLot{},
		strategy: constants.EvenDistribution,
	}, multiParking)
}

func TestPark_ShouldParkTheCar_GivenWeHaveEmptySlots(t *testing.T) {
	multiParking := New()
	parkingLot, _ := parkinglot.New(10)

	multiParking.AddParkingLot(parkingLot)

	car1 := models.NewCar("KA-01-HH-1234", "White")	
	slotNumber, err := multiParking.Park(car1)

	assert.Equal(t, 1, slotNumber)
	assert.NoError(t, err)
}
