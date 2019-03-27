package update

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Do() {
	url := "https://opendata.epa.gov.tw/ws/Data/AQI/?$format=json"

	var locationJSON []location

	err := getJSON(url, &locationJSON)
	check(err)

	content, err := json.MarshalIndent(locationJSON, "", "  ")
	check(err)

	publishTime := locationJSON[0].PublishTime
	t, err := time.Parse("2006-01-02 15:04", publishTime)

	savepath := fmt.Sprintf("json/%s/%s/%s.json", t.Format("2006-01"), t.Format("2006-01-02"), t.Format(time.RFC3339))

	// make dir
	err = os.MkdirAll(filepath.Dir(savepath), os.ModePerm)
	check(err)

	err = ioutil.WriteFile(savepath, content, 0644)
	check(err)

	fmt.Println("\ndone.")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getJSON(url string, target interface{}) error {

	r, err := http.Get(url)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	// data, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	return err
	// }

	return json.NewDecoder(r.Body).Decode(target)
}
