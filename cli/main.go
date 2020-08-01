package main

import (
	"log"

	"github.com/DACUS1995/anomaly-detector/pkg/dataset"
	"github.com/DACUS1995/anomaly-detector/pkg/detectors/external"
)

func main() {
	detector, err := external.NewBasicDetector("http://localhost")
	if err != nil {
		log.Fatal(err)
	}

	result := detector.Detect(&dataset.SimpleDatapoint{
		Values: []float32{0.75953972, 2.67458352, -2.19266376, -0.68166648, 0.94098628, -0.98737533, -1.58778538, 1.03491316, 0.94223391, -0.47893457, -0.7477086},
	})
	log.Print(result)
}
