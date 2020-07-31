package external

import "gopkg.in/errgo.v2/fmt/errors"

type BasicDetector struct {
	*RestClient
}

func NewBasicDetector(url string) (*BasicDetector, error) {
	client, err := NewRestClient(url)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	newBasicDetector := &BasicDetector{client}
	newBasicDetector.Init()
	return newBasicDetector, nil
}
