package fetcher

import (
	"encoding/json"
	"fmt"
	"github.com/margostino/openearth/cache"
	"github.com/margostino/openearth/graph/model"
)

func matchStringFor(expected *string, current string) bool {
	return expected == nil || (expected != nil && *expected == current)
}

func FetchDatasets(id *string, name *string, category *string) ([]*model.Dataset, error) {
	var datasets []*model.Dataset
	values := cache.GetData("datasets").([]interface{})
	for _, value := range values {
		dataset := model.Dataset{}
		data, _ := json.Marshal(value)
		fmt.Println(string(data))
		json.Unmarshal(data, &dataset)

		if matchStringFor(id, dataset.ID) && matchStringFor(name, dataset.Name) && matchStringFor(category, dataset.Category) {
			datasets = append(datasets, &dataset)
		}
	}
	return datasets, nil
}
