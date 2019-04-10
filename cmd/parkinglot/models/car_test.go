package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRegistrationNumber_ShouldReturnRegistrationNumberOfCar_WhenRegistrationNumberIsCalled(t *testing.T) {
	assert := assert.New(t)

	car := NewCar("MH-04-AY-1111", "Black")

	assert.Equal("MH-04-AY-1111", car.RegistrationNumber())
}

func TestColor_ShouldReturnColorOfCar_WhenColorIsCalled(t *testing.T) {
	assert := assert.New(t)

	car := NewCar("MH-04-AY-1111", "Black")

	assert.Equal("Black", car.Color())
}
