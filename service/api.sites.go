package service

import (
	"encoding/json"
	"net/http"

	"github.com/wowissu/AirWatcher/api"
)

func ApiSites() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		sitemap, _ := api.GetSites()

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Header().Set("Server", "A Go Web Server")
		w.WriteHeader(http.StatusOK)

		// TODO log Err
		outputByte, _ := json.Marshal(sitemap)
		w.Write(outputByte)
	}
}
