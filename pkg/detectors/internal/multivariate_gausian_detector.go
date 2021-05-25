package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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
	Threshold float32     `json:Threshold`
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


func (detector *MultivariateGaussianDetector) DetectProb(x *dataset.SimpleDatapoint) float32 {
	var detCovMatrix float64 = 0
	invConvMatrix := make([][]float64, len(detector.parameters.CovMatrix))
	k := len(detector.parameters.MeanArr)

	x_f64 := make([]float64, len(x.Features))
	for i, val := range x.Features {
		x_f64[i] = float64(val)
	}

	mean_f64 := make([]float64, len(detector.parameters.MeanArr))
	for i, val := range detector.parameters.MeanArr {
		mean_f64[i] = float64(val)
	}

	x_unsqeezed := [][]float64{x_f64}
	mean_unsqueezed := [][]float64{mean_f64}
	x_centered, _ := substract(x_unsqeezed, mean_unsqueezed)

	tmp, _ := multiply(transpose(x_centered), invConvMatrix)
	tmp, _ = multiply(tmp, x_centered)

	eExponent := -0.5 * tmp[0][0]
	fraction := 1 / math.Sqrt((math.Pow((2*math.Pi), float64(k)) * detCovMatrix))
	prob := fraction * math.Pow(math.E, eExponent)
}


func (detector *MultivariateGaussianDetector) Detect(x *dataset.SimpleDatapoint) bool {
	return detector.DetectProb(x) > detector.parameters.Threshold
}


func (detector *MultivariateGaussianDetector) Save() {
	detector.saveParamters(detector.parametersPath)
}


func (detector *MultivariateGaussianDetector) computeParameters(dataset dataset.Dataset) {
	sampleSize := len(dataset.GetSamples()[0].Features)

	// Init parameters
	meanArr := make([]float32, sampleSize)
	covMatrix := make([][]float32, sampleSize)
	threshold := float32(0)

	for idx := range covMatrix {
		covMatrix[idx] = make([]float32, sampleSize)
	}

	detector.parameters = &Parameters{
		MeanArr:   meanArr,
		CovMatrix: covMatrix,
		Threshold: threshold,
	}

	// Compute mean
	for _, elem := range dataset.GetSamples() {
		for idx, value := range elem.Features {
			meanArr[idx] += value
		}
	}

	// Compute covariance matrix
	for i := 0; i < sampleSize; i++ {
		for j := 0; j < sampleSize; j++ {
			var expectedVar float32 = 0
			for _, elem := range dataset.GetSamples() {
				firstElement := elem.Features[i] - meanArr[i]
				secondElement := elem.Features[j] - meanArr[j]
				expectedVar += firstElement * secondElement
			}
			covMatrix[i][j] = expectedVar / float32(sampleSize)
		}
	}

	// Compute the threshold https://towardsdatascience.com/wondering-how-to-build-an-anomaly-detection-model-87d28e50309
	threshold = detector.computeParameters(dataset)

	detector.parameters.MeanArr = meanArr
	detector.parameters.CovMatrix = covMatrix
}

// TODO compute tghe probs for each datapoint and then use them to pick the optimatl threshold based on the F1 score
func (detector *MultivariateGaussianDetector) computeThreshold(dataset dataset.Dataset) float32 {
	F1 := float32(0)

	min, max := MinMax()
	stepSize := 
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

func (detector *MultivariateGaussianDetector) saveParamters(savePath *string) error {
	jsonFile, err := os.Open(*savePath)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	defer jsonFile.Close()

	byteValue, err := json.Marshal(detector.parameters)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	err = ioutil.WriteFile(*savePath, byteValue, os.ModePerm)
	return err
}

func MinMax(array []float32) (float32, float32) {
	var max float32 = array[0]
	var min float32 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
