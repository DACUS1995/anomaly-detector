package external

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/DACUS1995/anomaly-detector/pkg/dataset"
	"github.com/DACUS1995/anomaly-detector/pkg/detectors"
	"gopkg.in/errgo.v2/fmt/errors"
)

type RestClient struct {
	url string
}

func NewRestClient(url string) (*RestClient, error) {
	NewRestClient := &RestClient{url: url}
	err := NewRestClient.Init()
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return NewRestClient, nil
}

func (client *RestClient) Init() error {
	return nil
}

func (client *RestClient) Detect(x *dataset.SimpleDatapoint) *detectors.Result {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(x)

	resp, err := http.Post(client.url+"/detect", "application/json", buf)
	if err != nil {
		// Do something
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Do something
	}

	var resultsMap map[string]interface{}
	json.Unmarshal([]byte(body), &resultsMap)

	return &detectors.Result{
		IsAnomaly: resultsMap["class_id"].(float64) == 1,
		Details:   resultsMap["class_name"].(string),
	}
}

func (client *RestClient) BatchDetect(x []dataset.SimpleDatapoint) []detectors.Result {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(x)

	resp, err := http.Post(client.url+"/detect/batch", "application/json", buf)
	if err != nil {
		// Do something
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Do something
	}

	var resultsMap []map[string]interface{}
	json.Unmarshal([]byte(body), &resultsMap)

	results := make([]detectors.Result, len(x))

	for idx, result := range resultsMap {
		results[idx] = detectors.Result{
			IsAnomaly: result["class_id"].(float64) == 1,
			Details:   result["class_name"].(string),
		}
	}

	return results
}
