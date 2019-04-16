package services

import (
	"fmt"
	"text/tabwriter"
	"os"
	"strings"
	"strconv"
	"parking_lot/cmd/parkinglot/parkinglot"
	"parking_lot/cmd/parkinglot/models"
	"parking_lot/cmd/parkinglot/constants"
)

type CommandService struct {
	parking *parkinglot.ParkingLot
}

func NewCommandService() *CommandService {
	return &CommandService{}
}

func (c *CommandService) Execute(command string) {
	commandParts := strings.Split(command, " ")

	switch commandParts[0] {
	case "create_parking_lot":
		numberOfSlots, err := strconv.Atoi(commandParts[1])
		if err != nil {
			fmt.Println(err)
		}
		c.createParkingLot(numberOfSlots)
	case "park":
		c.park(commandParts[1], commandParts[2])
	case "leave":
		slotNumber, err := strconv.Atoi(commandParts[1])
		if err != nil {
			fmt.Println(err)
		}
		c.leave(slotNumber)
	case "status":
		c.status()
	case "registration_numbers_for_cars_with_colour":
		c.registrationNumbersForCarsWithColour(commandParts[1])
	case "slot_numbers_for_cars_with_colour":
		c.slotNumbersForCarsWithColour(commandParts[1])
	case "slot_number_for_registration_number":
		c.slotNumberForRegistrationNumber(commandParts[1])
	default:
	}
}

func (c *CommandService) createParkingLot(noOfSlots int) {
	parking, err := parkinglot.New(noOfSlots)
	if err != nil {
		fmt.Println(err)
	} else {
		c.parking = parking
		fmt.Println(fmt.Sprintf(constants.CreateParkingLotSuccess, noOfSlots))
	}
}

func (c *CommandService) park(registrationNumber string, color string) {
	car := models.NewCar(registrationNumber, color)
	slotNumber, err := c.parking.Park(car)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf(constants.ParkSuccess, slotNumber))
	}
}

func (c *CommandService) leave(slotNumber int) {
	_, err := c.parking.UnPark(slotNumber)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf(constants.UnParkSuccess, slotNumber))
	}
}

func (c *CommandService) status() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 4, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "%s\t%s\t%s\n", constants.SlotNo, constants.RegistrationNo, constants.Color)

	details := c.parking.Status()
	for _, detail := range details {
		fmt.Fprintf(w, "%s\t%s\t%s\n", detail[0], detail[1], detail[2])
	}
}

func (c *CommandService) registrationNumbersForCarsWithColour(color string) {
	registrationNumbers, err := c.parking.RegistrationNumbersForCarsWithColor(color)
	if err != nil {
		fmt.Println(err)
	} else {
		for index, registrationNumber := range registrationNumbers {
			fmt.Print(registrationNumber)
			if index != len(registrationNumbers) - 1 {
				fmt.Print(", ")
			} 
		}
		fmt.Println()
	}
}

func (c *CommandService) slotNumbersForCarsWithColour(color string) {
	slotNumbers, err := c.parking.SlotNumbersForCarsWithColor(color)
	if err != nil {
		fmt.Println(err)
	} else {
		for index, slotNumber := range slotNumbers {
			fmt.Print(slotNumber)
			if index != len(slotNumbers) - 1 {
				fmt.Print(", ")
			} 
		}
		fmt.Println()
	}
}

func (c *CommandService) slotNumberForRegistrationNumber(registrationNumber string) {
	slotNumber, err := c.parking.SlotNumberForRegistrationNumber(registrationNumber)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(slotNumber)
	}
}
