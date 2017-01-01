package search

import (
	"os"
	"encoding/json"
)

type Feed struct {
	Name string `json:"site"`
	URI string `josn:"link"`
 	Type string `json:"type"`
}

const dataFile = "data/data.json"

func RetrieveFeeds() ([]*Feed, error)  {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
