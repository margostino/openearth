package fetcher

import (
	"encoding/json"
	"github.com/margostino/openearth/cache"
	"github.com/margostino/openearth/graph/model"
	"strings"
)

func matchStringFor(expected *string, current string) bool {
	return expected == nil || (expected != nil && *expected == current)
}

func FetchDatasets(id *string, name *string, category *string) ([]*model.Dataset, error) {
	var datasets []*model.Dataset
	values := cache.GetData(cache.Datasets).([]interface{})
	for _, value := range values {
		dataset := model.Dataset{}
		data, _ := json.Marshal(value)
		json.Unmarshal(data, &dataset)

		if matchStringFor(id, dataset.ID) && matchStringFor(name, dataset.Name) && matchStringFor(category, dataset.Category) {
			datasets = append(datasets, &dataset)
		}
	}
	return datasets, nil
}
func FetchNasaRssFeeds() ([]*model.NasaRssFeed, error) {
	var nasaRssFeeds []*model.NasaRssFeed
	data := cache.GetData(cache.NasaRssFeeds).([]interface{})
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &nasaRssFeeds)
	return nasaRssFeeds, nil
}

func FetchNasaEarthData(topicName *string) (*model.NasaEarthData, error) {
	var cachedNasaEarthData *model.NasaEarthData
	data := cache.GetData(cache.NasaEarthData).(map[string]interface{})
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &cachedNasaEarthData)

	var topics []*model.NasaEarthDataTopic

	if topicName == nil {
		topics = cachedNasaEarthData.Topics
	} else {
		topics = make([]*model.NasaEarthDataTopic, 0)
	}

	for _, topic := range cachedNasaEarthData.Topics {
		if topicName != nil && strings.ToLower(*topicName) == strings.ToLower(topic.Name) {
			topics = append(topics, topic)
		}
	}

	nasaEarthData := &model.NasaEarthData{
		URL:         cachedNasaEarthData.URL,
		Rss:         cachedNasaEarthData.Rss,
		Description: cachedNasaEarthData.Description,
		Topics:      topics,
	}

	//for _, value := range values {
	//	dataset := model.Dataset{}
	//	data, _ := json.Marshal(value)
	//	json.Unmarshal(data, &dataset)
	//
	//	//if matchStringFor(id, dataset.ID) && matchStringFor(name, dataset.Name) && matchStringFor(category, dataset.Category) {
	//	//	datasets = append(datasets, &dataset)
	//	//}
	//}
	return nasaEarthData, nil
}
