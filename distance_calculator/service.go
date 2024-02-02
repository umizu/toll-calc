package main

import (
	"math"

	"github.com/umizu/tollcalc/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculatorService struct {
	prevPoint []float64
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{}
}

func (c *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := 0.0
	if len(c.prevPoint) > 0 {
		distance = calculateDistance(c.prevPoint[0], c.prevPoint[1], data.Lat, data.Lon)
	}
	c.prevPoint = []float64{data.Lat, data.Lon}
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
