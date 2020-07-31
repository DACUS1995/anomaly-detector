package external

import (
	"bytes"
	"encoding/json"
	"net/http"

	"gopkg.in/errgo.v2/fmt/errors"
)

type PyodDetector struct {
	*RestClient
	algorithmType string
}

func NewPyodDetector(url string, algorithmType string) (*PyodDetector, error) {
	client, err := NewRestClient(url)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	newPyodDetector := &PyodDetector{client, algorithmType}
	newPyodDetector.Init()
	return newPyodDetector, nil
}

func (detector *PyodDetector) Init() error {
	buf := new(bytes.Buffer)

	initConfig := make(map[string]string)
	initConfig["type"] = detector.algorithmType
	json.NewEncoder(buf).Encode(initConfig)

	resp, err := http.Post(detector.url+"/init", "application/json", buf)
	if err != nil {
		return errors.Wrap(err)
	}
	defer resp.Body.Close()

	return nil
}
