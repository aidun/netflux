package netatmo

import (
	"time"
)

type Metric interface {
	GetTimestamp() *time.Time
	GetMetricName() string
	GetMetricValue() string
}

type NetatmoMetric struct {
}

func (m *NetatmoMetric) GetTimestamp() *time.Time {
	return nil
}

func (m *NetatmoMetric) GetMetricName() string {
	return "nil"
}

func (m *NetatmoMetric) GetMetricValue() string {
	return "nil"
}
