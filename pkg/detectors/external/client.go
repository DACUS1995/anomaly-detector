package external

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/DACUS1995/anomaly-detector/pkg/dataset"
	"github.com/DACUS1995/anomaly-detector/pkg/detectors"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{
		url: url,
	}
}

func (client *Client) Init() error {
	return nil
}

func (client *Client) Detect(x *dataset.SimpleDatapoint) *detectors.Result {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(x)

	resp, err := http.Post(client.url+"/detect", "image/jpeg", buf)
	if err != nil {
		// Do something
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Do something
	}

	var results map[string]interface{}
	json.Unmarshal([]byte(body), &results)

	return &detectors.Result{
		IsAnomaly: results["class_id"].(bool),
		Details:   "",
	}
}
