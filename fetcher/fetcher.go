package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/margostino/openearth/common"
	"github.com/margostino/openearth/metrics"
	"github.com/margostino/openearth/utils"
	"log"
	"strconv"
	"strings"
)

func Fetch(ctx context.Context, entity string, year int, response any) (map[string]*float64, error) {
	var results = make(map[string]*float64)
	dataset, err := getDatasetFrom(ctx)
	if err != nil {
		return nil, err
	}

	entity = strings.ToLower(entity)
	yearAsString := strconv.Itoa(year)
	dataKey := strings.ToLower(entity + yearAsString)

	log.Printf("Query for dataset %s with entity [%s] and year [%s]\n", dataset, entity, yearAsString)
	go metrics.PublishQuery(ctx)

	if value, ok := datasetCache[dataset]; ok {
		if result, ok := value.Index[dataKey]; ok {
			log.Println("Results from cache")
			results = result.Row
		}
	} else {
		url := indexCache[dataset]
		data, err := fetchCSVFromUrl(url)
		common.Check(err)

		var index = make(map[int]string)
		for idx, row := range data {
			row[0] = strings.ToLower(row[0])
			if idx == 0 {
				for dataIndex, column := range row[2:] {
					index[dataIndex] = utils.NormalizeName(strings.ToLower(column))
				}
				continue
			}

			if row[0] == entity && row[1] == yearAsString {
				for rowIndex, value := range row[2:] {
					resultValue, _ := strconv.ParseFloat(value, 8)
					results[index[rowIndex]] = &resultValue
				}

				datasetMapping := make(map[string]Dataset)
				cacheValue := Dataset{
					Row: results,
				}
				datasetMapping[dataKey] = cacheValue
				datasetIndex := DatasetIndex{
					Index: datasetMapping,
				}
				datasetCache[dataset] = datasetIndex
				log.Printf("New entry in cache for dataset %s and entity %s and year %s\n", dataset, entity, yearAsString)
			}
		}
	}

	// Convert map to json string
	jsonStr, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(jsonStr, response); err != nil {
		fmt.Println(err)
	}
	return results, nil
}
