package detector

import (
	"errors"
	"log"

	"github.com/DACUS1995/datacenter-anomaly-detector/pkg/dataset"
)

type MultivariateGaussianDetector struct {
	parametersPath *string
	datasetPath    *string
	meanArr        []float32
	covMatrix      [][]float32
}

type Parameters struct {
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
		detector.computeParameters(dataset)
	} else {
		detector.loadParameters()
	}
}

func (detector *MultivariateGaussianDetector) Detect(x *dataset.SimpleDatapoint) {

}

func (detector *MultivariateGaussianDetector) computeParameters(dataset dataset.Dataset) {
	sampleSize := len(dataset.GetSamples()[0].Values)
	meanArr := make([]float32, sampleSize)
	covMatrix := make([][]float32, sampleSize)
	for idx := range covMatrix {
		covMatrix[idx] = make([]float32, sampleSize)
	}

	// Compute mean
	for _, elem := range dataset.GetSamples() {
		for idx, value := range elem.Values {
			meanArr[idx] += value
		}
	}

	// Compute covariance matrix
	for i := 0; i < sampleSize; i++ {
		for j := 0; j < sampleSize; j++ {
			var expectedVar float32 = 0
			for _, elem := range dataset.GetSamples() {
				firstElement := elem.Values[i] - meanArr[i]
				secondElement := elem.Values[j] - meanArr[j]
				expectedVar += firstElement * secondElement
			}
			covMatrix[i][j] = expectedVar / float32(sampleSize)
		}
	}

	detector.meanArr = meanArr
	detector.covMatrix = covMatrix

}

func (detector *MultivariateGaussianDetector) loadParameters() {
	log.Println("Detector paramters loaded.")
}
