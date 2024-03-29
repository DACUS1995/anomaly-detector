package internal

import (
	"reflect"
	"testing"
)

func TestParametersLoad(t *testing.T) {
	parametersPath := "./params.json"

	detector := MultivariateGaussianDetector{
		parametersPath: &parametersPath,
		parameters:     &Parameters{},
	}

	detector.loadParameters()
	trueParams := Parameters{
		[]float32{0.1},
		[][]float32{{0.2}},
	}

	if !reflect.DeepEqual(trueParams.MeanArr, detector.parameters.MeanArr) {
		t.Errorf("Means array not equal")
	}

	if !reflect.DeepEqual(trueParams.CovMatrix, detector.parameters.CovMatrix) {
		t.Errorf("Covariance matrix not equal: %v", detector.parameters.CovMatrix)
	}
}

func TestParametersSave(t *testing.T) {
	parametersPath := "./params.json"

	detector := MultivariateGaussianDetector{
		parametersPath: &parametersPath,
		parameters:     &Parameters{},
	}
	detector.loadParameters()
	detector.Save()

	newDetector := MultivariateGaussianDetector{
		parametersPath: &parametersPath,
		parameters:     &Parameters{},
	}
	newDetector.loadParameters()

	if !reflect.DeepEqual(newDetector.parameters.MeanArr, detector.parameters.MeanArr) {
		t.Errorf("Means array not equal")
	}

	if !reflect.DeepEqual(newDetector.parameters.CovMatrix, detector.parameters.CovMatrix) {
		t.Errorf("Covariance matrix not equal: %v", detector.parameters.CovMatrix)
	}
}

func TestMultivariateGausianDetector(t *testing.T) {
	// parametersPath := "./params.json"
	// detector, err := NewMultivariateGaussianDetector(&parametersPath, nil)
	// if err != nil {
	// 	t.Error(err)
	// }

	// err = detector.Init()
	// if err != nil {
	// 	t.Error(err)
	// }
}
