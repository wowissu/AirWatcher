package main

import (
	"flag"

	"github.com/wowissu/AirWatcher/update"
)

func main() {
	isGetData := flag.Bool("getdata", false, "download json from https://opendata.epa.gov.tw")
	flag.Parse()

	// fmt.Printf("getdata: %t\n", *getdata)

	if *isGetData == true {
		update.Do()
	}
}
