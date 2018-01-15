package handler

import (
	"encoding/json"
	"net/http"

	"log"
)

func home(buildTime, commit, release string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		info := struct {
			BuildTime string `json:"buildTime"`
			Commit    string `json:"commit"`
			Release   string `json:"release"`
		}{
			buildTime, commit, release,
		}
		body, err := json.Marshal(info)
		if err != nil {
			log.Printf("不能编码信息数据：%v", err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
		log.Print("你好，你的请求已经被处理。")
	}
}
