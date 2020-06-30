package detector

import (
	"errors"
	"log"

	"github.com/DACUS1995/datacenter-anomaly-detector/pkg/dataset"
)

type MultivariateGaussianDetector struct {
	parametersPath *string
	datasetPath    *string
}

func NewMultivariateGaussianDetector(parametersPath *string, datasetPath *string) (*MultivariateGaussianDetector, error) {
	if parametersPath == nil {
		return &MultivariateGaussianDetector{
			datasetPath: datasetPath,
		}, nil
	} else {
		if datasetPath == nil {
			return nil, errors.New("Only one paramters can have a nil value.")
		}

		return &MultivariateGaussianDetector{
			parametersPath: parametersPath,
		}, nil
	}
}

func (detector *MultivariateGaussianDetector) Init() {
	if detector.parametersPath == nil {
		dataset := dataset.NewCSVDataset(*detector.datasetPath)
		detector.computeParamters(dataset)
	} else {
		detector.loadParameters()
	}
}

func (detector *MultivariateGaussianDetector) Detect(x *dataset.SimpleDatapoint) {

}

func (detector *MultivariateGaussianDetector) computeParamters(dataset dataset.Dataset) {
	sampleSize := len(dataset.GetSamples()[0].Values)

	meanArr := make([]float32, sampleSize)
	covMatrix := make([][]float32, sampleSize)
	for idx := range covMatrix {
		covMatrix[idx] = make([]float32, sampleSize)
	}

	for _, elem := range dataset.GetSamples() {
		for idx, value := range elem.Values {
			meanArr[idx] += value
		}
	}
}

func (detector *MultivariateGaussianDetector) loadParameters() {
	log.Println("Detector loaded.")
}
