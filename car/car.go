package car

type Wheel struct {
	Size     int
	Material string
}
type Car struct {
	Color       string
	Name        string
	WheelAmount int
	EngineType  string
	Wheel       Wheel
}

func (c *Car) GetColor() string {
	return c.Color
}

func (c *Car) setColor(color string) {
	c.Color = color
}

func (c *Car) GetName() string {
	return c.Name
}

func (c *Car) SetName(name string) {
	c.Name = name
}

func (c *Car) GetWheelAmount() int {
	return c.WheelAmount
}

func (c *Car) SetWheelAmount(wheelAmount int) {
	c.WheelAmount = wheelAmount
}

func (c *Car) GetEngineType() string {
	return c.EngineType
}

func (c *Car) SetEngineType(engineType string) {
	c.EngineType = engineType
}

//Embedded Struct getter and Setter

func (c *Car) GetWheel() *Wheel {
	return &c.Wheel
}

func (w *Wheel) GetWheelSize() int {
	return w.Size
}

func (w *Wheel) GetMaterial() string {
	return w.Material
}

func (w *Wheel) setWheelSize(wheelSize int) {
	w.Size = wheelSize
}
func (w *Wheel) setWheelMaterial(wheelMaterial string) {
	w.Material = wheelMaterial
}
