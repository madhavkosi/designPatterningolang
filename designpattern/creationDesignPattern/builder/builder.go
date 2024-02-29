package builder

type Car struct {
	Color         string
	EngineType    string
	HasSunroof    bool
	HasNavigation bool
}
type CarResponse struct {
	Car Car
}

func NewCarBuilder() *CarResponse {
	return &CarResponse{}
}

func (c *CarResponse) SetColor(color string) *CarResponse {
	c.Car.Color = color
	return c
}

func (c *CarResponse) SetEngineType(engineType string) *CarResponse {
	c.Car.EngineType = engineType
	return c
}

func (c *CarResponse) SetHasSunRoof(HasSunroof bool) *CarResponse {
	c.Car.HasSunroof = HasSunroof
	return c
}

func (c *CarResponse) SetHasNavigation(HasNavigation bool) *CarResponse {
	c.Car.HasNavigation = HasNavigation
	return c
}
func (c *CarResponse) Build() *Car {
	return &c.Car
}
