// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Dataset struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Source      *DatasetSource `json:"source"`
	URL         string         `json:"url"`
	LastUpdated string         `json:"last_updated"`
}

type DatasetSource struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}
