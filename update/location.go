package update

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type LocationUpdater struct{}

func (l *LocationUpdater) Update() error {
	// get sites.json
	sitesStorage := []SiteStorage{}
	siteStorageMap := make(map[string]*SiteStorage)

	sitesContent, err := ioutil.ReadFile(ApiDir("sites.json"))
	if err != nil {
		return err
	}
	json.Unmarshal(sitesContent, &sitesStorage)

	// gen sites name map
	for index, ss := range sitesStorage {
		siteStorageMap[ss.SiteName] = &sitesStorage[index]
	}

	// 今日
	t := time.Now()
	t = t.AddDate(0, 0, -1)

	hourlyFilesDir := filepath.Dir(new(SitesDataHourly).FileName(DataDir(), t))
	hourlyFiles := []string{}
	err = filepath.Walk(hourlyFilesDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			hourlyFiles = append(hourlyFiles, path)
		}
		return nil
	})
	check(err)

	for _, f := range hourlyFiles {
		hourSites := SitesDataHourly{}
		hourSites.Load(f)

		for _, s := range hourSites {
			// fmt.Println(s)
			siteStorageMap[s.SiteName].List = append(siteStorageMap[s.SiteName].List, s)
		}
	}

	for _, s := range siteStorageMap {
		fmt.Println(s.Dir)
		s.Save("today.json")
	}

	return nil
}
