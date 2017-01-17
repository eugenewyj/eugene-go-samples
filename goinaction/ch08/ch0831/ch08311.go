package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

type (
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL string `json:"unescapedUrl"`
		URL string `json:"url"`
		VisibleURL string `json:"visibleUrl"`
		CacheUrl string `json:"cacheUrl"`
		Title string `json:"title"`
		TitleNoFormatting string `json:"titleNoFormatting"`
		Content string `json:"content"`
	}

	gResponse struct {
		ResponseData struct{
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	uri := "https://ajax.googleapis.com/ajax/services/search/web?v=1.0&q=golang"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)

	pretty, err := json.MarshalIndent(gr, "", "    ")
	if err != nil {
		log.Println("ERROR:", err)
	}

	fmt.Println(string(pretty))
}
