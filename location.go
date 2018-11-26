package polis

type Location interface {
	Coordinates() Coordinates
	Distance(*Coordinates) float64
}
