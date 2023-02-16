package fetcher

import (
	"encoding/json"
	"fmt"
	"github.com/margostino/openearth/cache"
	"github.com/margostino/openearth/graph/model"
)

func FetchDatasets(id *string, name *string, category *string) ([]*model.Dataset, error) {
	var datasets []*model.Dataset
	values := cache.GetIndexBy("datasets").([]interface{})
	for _, value := range values {
		dataset := model.Dataset{}
		jsonString, _ := json.Marshal(value)
		fmt.Println(string(jsonString))
		json.Unmarshal(jsonString, &dataset)
		datasets = append(datasets, &dataset)
	}
	return datasets, nil
}

//func Fetch(ctx context.Context, entity string, year int, response any) (map[string]*float64, error) {
//	var results = make(map[string]*float64)
//	dataset, err := getDatasetFrom(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	entity = strings.ToLower(entity)
//	yearAsString := strconv.Itoa(year)
//	dataKey := strings.ToLower(entity + yearAsString)
//
//	log.Printf("Query for dataset %s with entity [%s] and year [%s]\n", dataset, entity, yearAsString)
//	go metrics.PublishQuery(ctx)
//
//	if value, ok := cache.datasetCache[dataset]; ok {
//		if result, ok := value.Index[dataKey]; ok {
//			log.Println("Results from cache")
//			results = result.Row
//		}
//	} else {
//		url := cache.indexCache[dataset]
//		data, err := fetchCSVFromUrl(url)
//		common.Check(err)
//
//		var index = make(map[int]string)
//		for idx, row := range data {
//			row[0] = strings.ToLower(row[0])
//			if idx == 0 {
//				for dataIndex, column := range row[2:] {
//					index[dataIndex] = utils.NormalizeName(strings.ToLower(column))
//				}
//				continue
//			}
//
//			if row[0] == entity && row[1] == yearAsString {
//				for rowIndex, value := range row[2:] {
//					resultValue, _ := strconv.ParseFloat(value, 8)
//					results[index[rowIndex]] = &resultValue
//				}
//
//				datasetMapping := make(map[string]Dataset)
//				cacheValue := Dataset{
//					Row: results,
//				}
//				datasetMapping[dataKey] = cacheValue
//				datasetIndex := DatasetIndex{
//					Index: datasetMapping,
//				}
//				cache.datasetCache[dataset] = datasetIndex
//				log.Printf("New entry in cache for dataset %s and entity %s and year %s\n", dataset, entity, yearAsString)
//			}
//		}
//	}
//
//	// Convert map to json string
//	jsonStr, err := json.Marshal(results)
//	if err != nil {
//		fmt.Println(err)
//	}
//	if err := json.Unmarshal(jsonStr, response); err != nil {
//		fmt.Println(err)
//	}
//	return results, nil
//}
