package multiparking

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"parking_lot/cmd/parkinglot/parkinglot"
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
