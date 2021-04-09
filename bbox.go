package bbox

import "errors"

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