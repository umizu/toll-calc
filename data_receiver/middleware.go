package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/umizu/tollcalc/types"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{next: next}
}

func (l *LogMiddleware) ProduceData(data types.OBUData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obu_id": data.OBUId,
			"lat":    data.Lat,
			"lon":    data.Lon,
			"took":   time.Since(start),
		}).Info("received data")
	}(time.Now())
	return l.next.ProduceData(data)
}
