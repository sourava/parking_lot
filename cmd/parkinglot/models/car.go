package models

type Car struct {
	registrationNumber string
	color string
}

func NewCar(registrationNumber string, color string) *Car {
	return &Car{
		registrationNumber: registrationNumber,
		color: color,
	}
}

func (c *Car) RegistrationNumber() string {
	return c.registrationNumber
}

func (c *Car) Color() string {
	return c.color
}
