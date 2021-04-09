# bbox

This is a Go library to get the bounding box from coordinates in floating-point or E7 represention.

## Install

````bash
go get github.com/lbrun25/bbox
````

## Usage

````go
package main

import (
	"bytes"
	"log"
	
	"github.com/lbrun25/fenv"
	"github.com/spf13/viper"
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