package bbox

import (
	"errors"
	"math"
)

var (
	// ErrEmptyCoordinates will throw if the coordinates are empty
	ErrEmptyCoordinates = errors.New("the coordinates are empty")
)

type BoundingBoxE7 struct {
	Left 	int32	`json:"left"`
	Bottom 	int32	`json:"bottom"`
	Right 	int32	`json:"right"`
	Top 	int32	`json:"top"`
}

// Coordinate are represented as latitude-longitude pairs in the E7 representation.
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees.
// Longitude should be in the range +/- 180 degrees (inclusive).
type CoordinateE7 struct {
	Lat int32 `json:"lat"`
	Lon	int32 `json:"lon"`
}

type BoundingBoxFloat struct {
	Left 	float64	`json:"left"`
	Bottom 	float64	`json:"bottom"`
	Right 	float64	`json:"right"`
	Top 	float64	`json:"top"`
}

// Coordinate are represented as latitude-longitude pairs in float.
// Latitudes should be in the range +/- 90 degrees.
// Longitude should be in the range +/- 180 degrees (inclusive).
type CoordinateFloat struct {
	Lat float64 `json:"lat"`
	Lon	float64 `json:"lon"`
}

// GetFromFloatCoordinates gets the bounding box from the floating-point representation of the latitude/longitude coordinates
func GetFromFloatCoordinates(coordinates []CoordinateFloat) (*BoundingBoxFloat, error) {
	if len(coordinates) == 0 {
		return nil, ErrEmptyCoordinates
	}
	minLat := coordinates[0].Lat
	minLon := coordinates[0].Lon
	maxLat := coordinates[0].Lat
	maxLon := coordinates[0].Lon

	for _, coordinate := range coordinates {
		if coordinate.Lat < minLat {
			minLat = coordinate.Lat
		}
		if coordinate.Lat > maxLat {
			maxLat = coordinate.Lat
		}
		if coordinate.Lon < minLon {
			minLon = coordinate.Lon
		}
		if coordinate.Lon > maxLon {
			maxLon = coordinate.Lon
		}
	}
	res := BoundingBoxFloat{
		Left:   minLat,
		Bottom: minLon,
		Right:  maxLat,
		Top:    maxLon,
	}
	return &res, nil
}

// GetFromE7Coordinates gets the bounding box from the integer representation of the latitude/longitude coordinates
func GetFromE7Coordinates(coordinates []CoordinateE7) (*BoundingBoxE7, error) {
	if len(coordinates) == 0 {
		return nil, ErrEmptyCoordinates
	}
	minLat := coordinates[0].Lat
	minLon := coordinates[0].Lon
	maxLat := coordinates[0].Lat
	maxLon := coordinates[0].Lon

	for _, coordinate := range coordinates {
		if coordinate.Lat < minLat {
			minLat = coordinate.Lat
		}
		if coordinate.Lat > maxLat {
			maxLat = coordinate.Lat
		}
		if coordinate.Lon < minLon {
			minLon = coordinate.Lon
		}
		if coordinate.Lon > maxLon {
			maxLon = coordinate.Lon
		}
	}
	res := BoundingBoxE7{
		Left:   minLat,
		Bottom: minLon,
		Right:  maxLat,
		Top:    maxLon,
	}
	return &res, nil
}

// getDistanceBBox gets the great circle distance
// Based on http://janmatuschek.de/LatitudeLongitudeBoundingCoordinates#Distance
func getDistanceBBox(left float64, bottom float64, right float64, top float64) float64 {
	lat1 := left * math.Pi / 180
	lon1 := bottom * math.Pi / 180
	lat2 := right * math.Pi / 180
	lon2 := top * math.Pi / 180

	r := 6371 * math.Pow10(3)
	dist := math.Acos(math.Sin(lat1) * math.Sin(lat2) + math.Cos(lat1) * math.Cos(lat2) * math.Cos(lon1 - lon2)) * r
	return dist
}

// GetFloatBBoxDistance gets the diagonal distance of the float bounding box
func GetFloatBBoxDistance(boundingBox BoundingBoxFloat) float64 {
	left := boundingBox.Left
	bottom := boundingBox.Bottom
	right := boundingBox.Right
	top := boundingBox.Top
	res := getDistanceBBox(left, bottom, right, top)
	return res
}

// GetE7BBoxDistance gets the diagonal distance of the E7 bounding box
func GetE7BBoxDistance(boundingBox BoundingBoxE7) float64 {
	left := float64(boundingBox.Left) * math.Pow10(-7)
	bottom := float64(boundingBox.Bottom) * math.Pow10(-7)
	right := float64(boundingBox.Right) * math.Pow10(-7)
	top := float64(boundingBox.Top) * math.Pow10(-7)
	res := getDistanceBBox(left, bottom, right, top)
	return res
}