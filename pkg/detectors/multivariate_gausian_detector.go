package detectors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/DACUS1995/anomaly-detector/pkg/dataset"
)

type MultivariateGaussianDetector struct {
	parametersPath *string
	datasetPath    *string
	parameters     *Parameters
}

type Parameters struct {
	MeanArr   []float32   `json:"MeanArr"`
	CovMatrix [][]float32 `json:"CovMatrix"`
}

func NewMultivariateGaussianDetector(parametersPath *string, datasetPath *string) (*MultivariateGaussianDetector, error) {
	if datasetPath == nil && parametersPath == nil {
		return nil, errors.New("Only one parameters can have a nil value")
	}

	if parametersPath == nil {
		return &MultivariateGaussianDetector{
			datasetPath: datasetPath,
			parameters:  &Parameters{},
		}, nil
	} else {
		return &MultivariateGaussianDetector{
			parametersPath: parametersPath,
			parameters:     &Parameters{},
		}, nil
	}
}

func (detector *MultivariateGaussianDetector) Init() error {
	if detector.parametersPath == nil && detector.datasetPath == nil {
		return errors.New("No parameters path or dataset path was set")
	}

	if detector.parametersPath == nil {
		dataset := dataset.NewCSVDataset(*detector.datasetPath)
		detector.computeParameters(dataset)
	} else {
		detector.loadParameters()
	}

	return nil
}

func (detector *MultivariateGaussianDetector) Detect(x *dataset.SimpleDatapoint) {

}

func (detector *MultivariateGaussianDetector) Save() {

}

func (detector *MultivariateGaussianDetector) computeParameters(dataset dataset.Dataset) {
	sampleSize := len(dataset.GetSamples()[0].Values)

	// Init parameters
	meanArr := make([]float32, sampleSize)
	covMatrix := make([][]float32, sampleSize)

	for idx := range covMatrix {
		covMatrix[idx] = make([]float32, sampleSize)
	}

	detector.parameters = &Parameters{
		MeanArr:   meanArr,
		CovMatrix: covMatrix,
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

	detector.parameters.MeanArr = meanArr
	detector.parameters.CovMatrix = covMatrix
}

func (detector *MultivariateGaussianDetector) loadParameters() {
	jsonFile, err := os.Open(*detector.parametersPath)
	if err != nil {
		fmt.Printf("Precomputed parameters file not found %v", *detector.parametersPath)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &detector.parameters)

	log.Println("Detector loaded.")
}
