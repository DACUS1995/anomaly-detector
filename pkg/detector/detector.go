package detector

import (
	"github.com/DACUS1995/datacenter-anomaly-detector/pkg/dataset"
)

type Detector interface {
	Init()
	Detect(x *dataset.SimpleDatapoint) *Result
}

type Result struct {
	IsAnomaly bool
	Details   string
}
