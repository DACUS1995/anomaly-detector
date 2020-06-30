package dataset

type CSVDataset struct {
	Samples []SimpleDatapoint
}

func NewCSVDataset(path string) *CSVDataset {
	dataset := &CSVDataset{}
	dataset.Load(path)
	return dataset
}

func (dataset *CSVDataset) GetSamples() []SimpleDatapoint {
	return dataset.Samples
}

func (dataset *CSVDataset) Load(path string) {

}

func (dataset *CSVDataset) Preprocess() {

}
