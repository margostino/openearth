package cache

import (
	"fmt"
	"github.com/margostino/openearth/common"
	"github.com/margostino/openearth/graph/model"
	"io"
	"net/http"
	"os"
)

var indexes = map[string]any{
	"datasets":  &model.Dataset{},
	"companies": &model.DatasetSource{},
}

var index = load()

func load() map[string]interface{} {
	var cache = make(map[string]any)
	baseUrl := os.Getenv("DATA_URL")

	for key, value := range indexes {
		url := fmt.Sprintf("%s/%s.yml", baseUrl, key)
		resp, err := http.Get(url)

		if !common.IsError(err, fmt.Sprintf("when fetching data for %s", key)) && resp.StatusCode == 200 {
			bytes, err := io.ReadAll(resp.Body)
			if !common.IsError(err, fmt.Sprintf("when reading data response for %s", key)) {
				common.UnmarshalYamlBytes(bytes, &value)
				cache[key] = value
			}
			resp.Body.Close()
		}

	}

	return cache
}

func GetIndexBy(key string) interface{} {
	return index[key]
}
