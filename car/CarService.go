package car

import "fmt"

type CarService struct {
}

func NewCarService() *CarService {
	return &CarService{}
}

func (cs *CarService) CreateCar(spec Car) (*Car, error) {
	if spec.WheelAmount < 2 {
		return nil, fmt.Errorf("cars must have 2 wheels")
	}

	car := &Car{
		Name:        spec.Name,
		Color:       spec.Color,
		EngineType:  spec.EngineType,
		WheelAmount: spec.WheelAmount,
		Wheel:       spec.Wheel,
	}
	return car, nil
}
