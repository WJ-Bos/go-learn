package car

import "fmt"

type CarService struct {
}

type ValidatedFields struct {
	Color   string
	Name    string
	WheelAm int
	EngT    string
	Wheel   Wheel
}

func Construct() *CarService {
	return &CarService{}
}

func (cs *CarService) NewCarBySpec(spec Car) (*Car, error) {
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

func (cs *CarService) NewCarEmpty() *Car {
	return &Car{}
}

func (cs *CarService) NewCarByValue(
	c string, n string, wheelAm int, engT string, wheel Wheel,
) (*Car, error) {
	validated, err := validateFields(c, n, wheelAm, engT, wheel)
	if err != nil {
		return &Car{
			Color:       c,
			Name:        n,
			WheelAmount: wheelAm,
			EngineType:  engT,
			Wheel:       wheel,
		}, fmt.Errorf("validation applied defaults: %w", err)
	}
	return &Car{
		Color:       validated.Color,
		Name:        validated.Name,
		WheelAmount: validated.WheelAm,
		EngineType:  validated.EngT,
		Wheel:       validated.Wheel,
	}, nil
}

func validateFields(
	c string, n string, wheelAm int, engT string, wheel Wheel,
) (ValidatedFields, error) {
	var errs []error
	vf := ValidatedFields{}
	if c == "" {
		vf.Color = "black"
	}
	if n == "" {
		vf.Name = "N/A"
	}
	if wheelAm < 2 {
		vf.WheelAm = 2
		errs = append(errs, fmt.Errorf("cars must have at least 2 wheels, defaulted to 2"))
	}
	if engT == "" {
		vf.EngT = "N/A"
	}
	if wheel == (Wheel{}) {
		errs = append(errs, fmt.Errorf("wheel struct cannot be empty"))
		vf.Wheel = Wheel{Size: 30, Material: "Rubber"}
	}
	if wheel.Size < 10 {
		vf.Wheel.Size = 30
		errs = append(errs, fmt.Errorf("wheel size must be at least 30"))
	}
	if wheel.Material == "" {
		vf.Wheel.Material = "Rubber"
	}

	var err error
	if len(errs) > 0 {
		err = fmt.Errorf("validation errors: %v", errs)
	}

	return vf, err
}
