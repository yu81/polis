package polis

import "math"

const (
	RadiusOfEarthKm = 6371.0
	// Spherical Angle in degree.
	LongitudeMax = 180.0
	LangitudeMin = -180.0
	LatitudeMax  = 90.0
	LatitudeMin  = -90.0
)

type Location struct {
	Lon float64 `json:"lon"` // degree
	Lat float64 `json:"lat"` // degree
}

func NewLocation(lonDegree, latDegree float64) Location {
	return Location{Lon: lonDegree, Lat: latDegree}
}

func (l *Location) Validate() bool {
	return l.Lon >= LangitudeMin && l.Lon <= LongitudeMax && l.Lat >= LatitudeMin && l.Lat <= LatitudeMax
}

func (l *Location) Longitude() float64 {
	return l.Lon
}

func (l *Location) Latitude() float64 {
	return l.Lat
}

func (l *Location) Radians() (float64, float64) {
	return math.Pi / 180.0 * l.Lon, math.Pi / 180.0 * l.Lat
}

func (l *Location) Distance(l2 *Location) float64 {
	lon1, lat1 := l.Radians()
	lon2, lat2 := l2.Radians()
	return Distance(lon1, lat1, lon2, lat2)
}

func Distance(lon1, lat1, lon2, lat2 float64) float64 {
	sp1, sp2 := math.Sin(lat1), math.Sin(lat2)
	cp1, cp2 := math.Cos(lat1), math.Cos(lat2)
	deltaLambda := lon1 - lon2
	if deltaLambda < 0 || deltaLambda > 180.0 {
		deltaLambda = lon2 - lon1
	}
	sdl, cdl := math.Sin(deltaLambda), math.Cos(deltaLambda)
	firstTerm, secondTerm := (cp2 * sdl), (cp1*sp2 - sp1*cp2*cdl)
	deltaSigma := math.Atan2(
		math.Sqrt(firstTerm*firstTerm+secondTerm*secondTerm),
		sp1*sp2+cp1*cp2*cdl)
	return RadiusOfEarthKm * deltaSigma
}
