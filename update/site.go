package update

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

// SiteDataHourly ...
type SiteDataHourly struct {
	SiteName    string `json:"SiteName"`
	County      string `json:"County"`
	AQI         string `json:"AQI"`
	Pollutant   string `json:"Pollutant"`
	Status      string `json:"Status"`
	SO2         string `json:"SO2"`
	CO          string `json:"CO"`
	CO8hr       string `json:"CO_8hr"`
	O3          string `json:"O3"`
	O38hr       string `json:"O3_8hr"`
	PM10        string `json:"PM10"`
	PM25        string `json:"PM2.5"`
	NO2         string `json:"NO2"`
	NOx         string `json:"NOx"`
	NO          string `json:"NO"`
	WindSpeed   string `json:"WindSpeed"`
	WindDirec   string `json:"WindDirec"`
	PublishTime string `json:"PublishTime"`
	PM25AVG     string `json:"PM2.5_AVG"`
	PM10AVG     string `json:"PM10_AVG"`
	SO2AVG      string `json:"SO2_AVG"`
	Longitude   string `json:"Longitude"`
	Latitude    string `json:"Latitude"`
}

// SitesDataHourly ...
type SitesDataHourly []SiteDataHourly

// Load ...
func (s *SitesDataHourly) Load(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		r, err := http.Get(path)
		if err != nil {
			return err
		}

		defer r.Body.Close()
		return json.NewDecoder(r.Body).Decode(s)
	} else {
		// fmt.Println("is exist file ...")
		file := path

		content, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		return json.Unmarshal(content, &s)
	}
}

// Save ...
func (s SitesDataHourly) Save(dir string) error {
	fmt.Println("json.MarshalIndent")
	content, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02 15:04", s[0].PublishTime)
	savepath := s.FileName(dir, t)
	savedir := filepath.Dir(savepath)

	// make dir
	if _, err := os.Stat(savedir); os.IsNotExist(err) {
		fmt.Println("MKdir...", savedir)
		err = os.MkdirAll(savedir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	fmt.Println("write file...", savepath)
	err = ioutil.WriteFile(savepath, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

// FileName ...
func (s *SitesDataHourly) FileName(dir string, t time.Time) string {
	return fmt.Sprintf("%s/%s/%s/%s.json", dir, t.Format("2006-01"), t.Format("2006-01-02"), t.Format("2006-01-02T15:00:00Z"))
}

// Dirs ...
func (s SitesDataHourly) Storages() ([]SiteStorage, error) {
	var storages []SiteStorage

	if len(s) <= 0 {
		return storages, errors.New("Do SitesDataHourly.Load(uri string) first")
	}

	for i, site := range s {
		id := i + 1

		storages = append(storages, SiteStorage{
			ID:       id,
			SiteName: site.SiteName,
			County:   site.County,
			Dir:      ApiDir("sites", strconv.Itoa(id)),
		})
	}

	return storages, nil
}

// SiteStorage ...
type SiteStorage struct {
	ID       int             `json:"id"`
	SiteName string          `json:"SiteName"`
	County   string          `json:"County"`
	Dir      string          `json:"Dir"`
	List     SitesDataHourly `json:"List"`
}

func (s *SiteStorage) Load(file string) error {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, &s.List)
}

// Save site storage in today or thisWeek or thisMonth
func (s *SiteStorage) Save(filename string) error {
	var content []byte
	var err error
	if content, err = json.MarshalIndent(s.List, "", "  "); err != nil {
		return err
	}

	// make dir
	if _, err := os.Stat(s.Dir); os.IsNotExist(err) {
		fmt.Println("MKdir...", s.Dir)
		err = os.MkdirAll(s.Dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	savepath := path.Join(s.Dir, filename)

	fmt.Println("write file...", savepath)
	if err := ioutil.WriteFile(savepath, content, 0644); err != nil {
		return err
	}

	return nil
}
