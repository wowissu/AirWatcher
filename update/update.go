package update

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

func Do() {
	url := "https://opendata.epa.gov.tw/ws/Data/AQI?$format=json"

	var locationJSON []location

	err := getJSON(url, &locationJSON)
	check(err)

	fmt.Println("json.MarshalIndent")
	content, err := json.MarshalIndent(locationJSON, "", "  ")
	check(err)

	t, err := time.Parse("2006-01-02 15:04", locationJSON[0].PublishTime)
	dir, err := os.Getwd()
	savepath := fmt.Sprintf("%s/json/%s/%s/%s.json", dir, t.Format("2006-01"), t.Format("2006-01-02"), t.Format(time.RFC3339))
	savedir := filepath.Dir(savepath)

	// make dir
	if _, err := os.Stat(savedir); os.IsNotExist(err) {
		fmt.Println("MKdir...", savedir)
		err = os.MkdirAll(savedir, os.ModePerm)
		check(err)
	}

	fmt.Println("write file...", savepath)
	err = ioutil.WriteFile(savepath, content, 0644)
	check(err)

	fmt.Println("done.\n")
}

//Log checks and logs a error
func check(err error) {
	if err != nil {
		log.Printf("%+v", errors.Wrap(err, ""))
		os.Exit(1)
	}
}

func getJSON(uri string, target interface{}) error {
	fmt.Println("Get json...", uri)
	r, err := http.Get(uri)

	if err != nil {
		return err
	}

	defer r.Body.Close()
	fmt.Println("Decode")
	return json.NewDecoder(r.Body).Decode(target)
}
