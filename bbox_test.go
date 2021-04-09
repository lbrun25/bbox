package bbox_test

import (
	"github.com/lbrun25/bbox"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestGetFromFloatCoordinates(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		coordinates := []bbox.CoordinateFloat{
			{
				// Left
				Lat: 48.581635,
				Lon: 7.744521,
			},
			{
				// Bottom
				Lat: 48.578911,
				Lon: 7.748500,
			},
			{
				// Right
				Lat: 48.581728,
				Lon: 7.764263,
			},
			{
				// Top
				Lat: 48.584880,
				Lon: 7.762572,
			},
			{
				// Middle
				Lat: 48.581851,
				Lon: 7.750945,
			},
		}

		// http://bboxfinder.com/#48.578911,7.744521,48.58488,7.764263
		expectedBbox := &bbox.BoundingBoxFloat{
			Left:   48.578911,
			Bottom: 7.744521,
			Right:  48.58488,
			Top:    7.764263,
		}
		actualBbox, err := bbox.GetFromFloatCoordinates(coordinates)
		assert.NoError(t, err)
		assert.Equal(t, expectedBbox, actualBbox)
	})
	
	t.Run("error: empty coordinates", func(t *testing.T) {
		var coordinates []bbox.CoordinateFloat

		actualBbox, err := bbox.GetFromFloatCoordinates(coordinates)
		assert.Error(t, err, bbox.ErrEmptyCoordinates)
		assert.Nil(t, actualBbox)
	})
}

func TestGetFromE7Coordinates(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		coordinates := []bbox.CoordinateE7{
			{
				// Left
				Lat: int32(48.581635 * math.Pow10(7)),
				Lon:  int32(7.744521 * math.Pow10(7)),
			},
			{
				// Bottom
				Lat: int32(48.578911 * math.Pow10(7)),
				Lon: int32(7.748500 * math.Pow10(7)),
			},
			{
				// Right
				Lat: int32(48.581728 * math.Pow10(7)),
				Lon: int32(7.764263 * math.Pow10(7)),
			},
			{
				// Top
				Lat: int32(48.584880 * math.Pow10(7)),
				Lon: int32(7.762572 * math.Pow10(7)),
			},
			{
				// Middle
				Lat: int32(48.581851 * math.Pow10(7)),
				Lon: int32(7.750945 * math.Pow10(7)),
			},
		}

		// http://bboxfinder.com/#48.578911,7.744521,48.58488,7.764263
		expectedBbox := &bbox.BoundingBoxE7{
			Left:   int32(48.578911 * math.Pow10(7)),
			Bottom: int32(7.744521 * math.Pow10(7)),
			Right:  int32(48.58488 * math.Pow10(7)),
			Top:    int32(7.764263 * math.Pow10(7)),
		}
		actualBbox, err := bbox.GetFromE7Coordinates(coordinates)
		assert.NoError(t, err)
		assert.Equal(t, expectedBbox, actualBbox)
	})

	t.Run("error: empty coordinates", func(t *testing.T) {
		var coordinates []bbox.CoordinateE7

		actualBbox, err := bbox.GetFromE7Coordinates(coordinates)
		assert.Error(t, err, bbox.ErrEmptyCoordinates)
		assert.Nil(t, actualBbox)
	})
}