package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/umizu/tollcalc/types"
)

type LogMiddleware struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) CalculatorServicer {
	return &LogMiddleware{next: next}
}

func (m *LogMiddleware) CalculateDistance(data types.OBUData) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"dist": dist,
			"err":  err,
		}).Info("calculated distance")
	}(time.Now())

	return m.next.CalculateDistance(data)
}
