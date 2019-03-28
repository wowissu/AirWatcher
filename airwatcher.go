package main

import (
	"flag"

	"github.com/wowissu/AirWatcher/update"
)

func main() {
	doUpdate := flag.Bool("update", false, "download json from https://opendata.epa.gov.tw")
	flag.Parse()

	// fmt.Printf("getdata: %t\n", *getdata)

	if *doUpdate == true {
		update.Do()
	}
}
