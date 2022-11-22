package nasa

import (
	"github.com/margostino/earth-station-api/server"
	"github.com/mmcdole/gofeed"
	"log"
	"net/http"
)

func GetTopics(feedUrls map[string]string, writer http.ResponseWriter) {
	topics := make([]string, 0)
	for topic, _ := range feedUrls {
		topics = append(topics, topic)
	}
	server.WriteResponse(topics, writer)
}

func GetTopic(topic string, feedUrls map[string]string, writer http.ResponseWriter) {

	items := make([]map[string]interface{}, 0)
	feedUrl := feedUrls[topic]
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedUrl)
	if feed != nil {
		for _, entry := range feed.Items {
			item := make(map[string]interface{})
			item["title"] = entry.Title
			item["description"] = entry.Description
			item["content"] = entry.Content
			item["link"] = entry.Link
			item["updated"] = entry.Updated
			item["published"] = entry.PublishedParsed.String()
			item["authors"] = entry.Authors
			items = append(items, item)
		}
	} else {
		log.Printf("There are no feeds")
	}

	server.WriteResponse(items, writer)

}
