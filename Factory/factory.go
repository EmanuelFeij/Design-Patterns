package Factory

import "fmt"

type MotorVehicle interface {
	Start() (bool, error)
}

type Car struct{}

func (c Car) Start() (bool, error) {
	fmt.Println("Car started")
	return true, nil
}

type Bike struct{}

func (b Bike) Start() (bool, error) {
	fmt.Println("Bike started")
	return true, nil
}

type Plane struct{}

func (p Plane) Start() (bool, error) {
	fmt.Println("Plane started")
	return true, nil
}

func FactoryOfMotorVehicle(vehicle string) MotorVehicle {
	var m MotorVehicle
	switch vehicle {
	case "car":
		m = Car{}
	case "bike":
		m = Bike{}
	case "plane":
		m = Plane{}
	default:
		m = nil
	}
	return m
}
