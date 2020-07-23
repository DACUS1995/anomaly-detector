package detectors

import (
	"github.com/DACUS1995/anomaly-detector/pkg/dataset"
)

type Detector interface {
	Init() error
	Detect(x *dataset.SimpleDatapoint) *Result
	Save(string)
}

type Result struct {
	IsAnomaly bool
	Details   string
}
