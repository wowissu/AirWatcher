package update

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
)

// Dir Save dir
var Dir string

func DataDir(suffix ...string) string {
	suffix = append([]string{Dir, "json/data"}, suffix...)
	return path.Join(suffix...)
}

func ApiDir(suffix ...string) string {
	suffix = append([]string{Dir, "json/api"}, suffix...)
	return path.Join(suffix...)
}

// Do the basic update
func Do() {
	url := "http://opendata.epa.gov.tw/ws/Data/AQI?$format=json"

	sites := new(SitesDataHourly)
	check(sites.Load(url))
	check(sites.Save(DataDir()))

	fmt.Println("done.")
}

// Location json update
func Location() {
	searchDir := DataDir()
	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	check(err)

	lastFile := fileList[len(fileList)-1]

	// SitesDataHourly Load
	var sitedataHours = new(SitesDataHourly)
	err = sitedataHours.Load(lastFile)
	check(err)

	// Gen []SiteMap
	siteStorages, err := sitedataHours.Storages()
	check(err)

	// encode json
	content, err := json.MarshalIndent(siteStorages, "", "  ")
	check(err)

	// check save path
	savepath := ApiDir("sites.json")
	savedir := filepath.Dir(savepath)
	err = mkdir(savedir)
	check(err)

	// write json into file
	fmt.Println("write file...", savepath)
	err = ioutil.WriteFile(savepath, content, 0644)
	check(err)
}

// Log checks and logs a error
func check(err error) {
	if err != nil {
		log.Printf("%+v", errors.Wrap(err, ""))
	}
}

func mkdir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}
