package polis

type City struct {
	Name  string      `json:"name"`
	Coord Coordinates `json:"coordinates"`
}

var _ = Location(&City{})

func (c *City) Coordinates() Coordinates {
	return c.Coord
}

func (c *City) Distance(l *Coordinates) float64 {
	return c.Coord.Distance(l)
}

func (c *City) NextTo() []Location {
	return []Location{}
}
