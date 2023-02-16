package cache

import (
	"fmt"
	"github.com/margostino/openearth/common"
	"github.com/margostino/openearth/config"
	"github.com/margostino/openearth/graph/model"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
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
	var bytes []byte
	var err error
	baseDataPath := os.Getenv("DATA_PATH")

	for key, value := range indexes {
		location := fmt.Sprintf("%s/%s.yml", baseDataPath, key)

		if config.IsDevEnv() {
			bytes, err = ioutil.ReadFile(location)
		} else {
			resp, err := http.Get(location)
			if !common.IsError(err, fmt.Sprintf("when fetching data for %s", key)) && resp.StatusCode == 200 {
				bytes, err = io.ReadAll(resp.Body)
				resp.Body.Close()
			}
		}

		if !common.IsError(err, fmt.Sprintf("when reading data for %s", key)) && bytes != nil {
			err = yaml.Unmarshal(bytes, &value)
			if !common.IsError(err, "when unmarshaling YAML data") {
				cache[key] = value
			}
		}

	}

	return cache
}

func GetData(key string) interface{} {
	return index[key]
}
