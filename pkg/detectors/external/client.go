package external

import (
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
	return nil
}
