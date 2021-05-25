package dataset

type Dataset interface {
	Load(string)
	Preprocess()
	GetSamples() []SimpleDatapoint
}

type SimpleDatapoint struct {
	Features []float32 `json:"features"`
	Label    []float32 `json:"label"`
}

type BatchDatapoint struct {
	Features [][]float32 `json:"features"`
	Labels   [][]float32 `json:"labels"`
}
