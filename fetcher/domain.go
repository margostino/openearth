package fetcher

type Dataset struct {
	Row map[string]*float64
}

type DatasetIndex struct {
	Index map[string]Dataset
}
