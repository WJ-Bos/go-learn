package main

import (
	"HelloGo/car"
	"fmt"
)

func main() {
	carService := car.NewCarService()

	newCar, err := carService.CreateCar(car.Car{
		Name:        "Tesla",
		Color:       "red",
		EngineType:  "Electric",
		WheelAmount: 4,
		Wheel: car.Wheel{
			Size:     40,
			Material: "Matte",
		},
	})
	if err != nil {
		fmt.Errorf("%v", err)
	}

	fmt.Printf("Created car: %+v\n", newCar)
}
