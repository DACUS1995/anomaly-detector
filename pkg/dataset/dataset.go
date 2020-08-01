package dataset

type Dataset interface {
	Load(string)
	Preprocess()
	GetSamples() []SimpleDatapoint
}

type SimpleDatapoint struct {
	Values []float32 `json:"values"`
}

type BatchDatapoint struct {
	Values [][]float32 `json:"values"`
}
