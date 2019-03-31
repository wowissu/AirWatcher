package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/wowissu/AirWatcher/update"
)

var Dir string

func GetSites() ([]update.SiteStorage, error) {
	siteStorages := []update.SiteStorage{}
	filepath := update.ApiDir("sites.json")

	fmt.Println(filepath)

	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return siteStorages, err
	}

	json.Unmarshal(content, &siteStorages)

	return siteStorages, nil
}
