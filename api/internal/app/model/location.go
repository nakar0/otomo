package model

type Location struct {
	GooglePlaceID string
	LongName      string
	ShortName     string
	Address       string
	Types         []string
	Geometry      Geometry
}

type Geometry struct {
	LatLng LatLng
}

type LatLng struct {
	Lat float64
	Lng float64
}
