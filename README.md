# bbox

This is a Go library to get the bounding box from coordinates in floating-point or E7 representation.

## Install

````bash
go get github.com/lbrun25/bbox
````

## Usage

````go
package main

import (
	"github.com/lbrun25/bbox"
)

func main() {
	coordinates := []bbox.CoordinateFloat{
		{
			Lat: 48.581635,
			Lon: 7.744521,
		},
		{
			Lat: 48.578911,
			Lon: 7.748500,
		},
		{
			Lat: 48.581728,
			Lon: 7.764263,
		},
		{
			Lat: 48.584880,
			Lon: 7.762572,
		},
		{
			Lat: 48.581851,
			Lon: 7.750945,
		},
	}
	boundingBox, err := bbox.GetFromFloatCoordinates(coordinates)
}
````
