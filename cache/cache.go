package cache

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/margostino/openearth/common"
	"github.com/margostino/openearth/config"
	"github.com/margostino/openearth/graph/model"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

const (
	RepoOwner = "margostino"
	RepoName  = "data"
)

var indexes = map[string]any{
	Datasets:      &model.Dataset{},
	NasaEarthData: &model.NasaEarthData{},
	NasaRssFeeds:  &model.NasaRssFeed{},
}

var index = load()
var githubClient = getGithubClient()

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
			bytes, err = getDataBy(location)
		}

		if !common.IsError(err, fmt.Sprintf("when reading data for %s", key)) && bytes != nil {
			err = yaml.Unmarshal(bytes, &value)
			if !common.IsError(err, "when unmarshalling YAML data") {
				cache[key] = value
			}
		}

	}

	return cache
}

func GetData(key string) interface{} {
	return index[key]
}

func getGithubClient() *github.Client {
	if !config.IsDevEnv() {
		var githubAccessToken = os.Getenv("GITHUB_ACCESS_TOKEN")
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: githubAccessToken},
		)
		tc := oauth2.NewClient(ctx, ts)
		return github.NewClient(tc)
	}
	return nil
}

func getDataBy(path string) ([]byte, error) {
	var bytes []byte
	options := &github.RepositoryContentGetOptions{}

	encodedContent, _, response, err := githubClient.Repositories.GetContents(context.Background(), RepoOwner, RepoName, path, options)
	if !common.IsError(err, "when getting data content from repository") && response.StatusCode == 200 {
		data, err := encodedContent.GetContent()
		return []byte(data), err
	}

	return bytes, err
}
