package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuadrantDefiner1Quadrant(t *testing.T) {

	quadrant1 := QuadrantDetector(Point{
		X: 1.5,
		Y: 1.9,
	})
	assert.Equal(t, 1, quadrant1)

	quadrant2 := QuadrantDetector(Point{
		X: 345.234,
		Y: 2345.654,
	})
	assert.Equal(t, 1, quadrant2)
}

func TestQuadrantDefiner2Quadrant(t *testing.T) {

	quadrant1 := QuadrantDetector(Point{
		X: -23.5,
		Y: 154.43,
	})
	assert.Equal(t, 2, quadrant1)

	quadrant2 := QuadrantDetector(Point{
		X: -345.323,
		Y: 2345.6,
	})
	assert.Equal(t, 2, quadrant2)
}

func TestQuadrantDefiner3Quadrant(t *testing.T) {

	quadrant1 := QuadrantDetector(Point{
		X: -11.5,
		Y: -154.34,
	})
	assert.Equal(t, 3, quadrant1)

	quadrant2 := QuadrantDetector(Point{
		X: -54.01,
		Y: -98.6,
	})
	assert.Equal(t, 3, quadrant2)
}

func TestQuadrantDefiner4Quadrant(t *testing.T) {

	quadrant1 := QuadrantDetector(Point{
		X: 11.5,
		Y: -154.34,
	})
	assert.Equal(t, 4, quadrant1)

	quadrant2 := QuadrantDetector(Point{
		X: 91.22,
		Y: -41.43,
	})
	assert.Equal(t, 4, quadrant2)
}

func TestQuadrantDefinerInvalidQuadrant(t *testing.T) {

	quadrant1 := QuadrantDetector(Point{
		X: 0,
		Y: 0,
	})
	assert.Equal(t, -1, quadrant1)
}

func TestIsInRadius(t *testing.T) {

	inRadius1 := IsInRadius(10, Point{
		X: 2.0,
		Y: 2.0,
	})
	assert.True(t, inRadius1)

	inRadius2 := IsInRadius(2, Point{
		X: 2.0,
		Y: 2.0,
	})
	assert.False(t, inRadius2)

	inRadius3 := IsInRadius(1, Point{
		X: 0.3,
		Y: 0.4,
	})
	assert.True(t, inRadius3)
}
