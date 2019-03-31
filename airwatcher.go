package main

import (
	"flag"
	"path"
	"runtime"

	"github.com/wowissu/AirWatcher/service"
	"github.com/wowissu/AirWatcher/update"
)

var Dir string

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	Dir = path.Dir(filename)
	update.Dir = Dir

	doUpdate := flag.Bool("update", false, "download json from https://opendata.epa.gov.tw")
	doUpdateLocation := flag.Bool("location", false, "update location.json")
	runApp := flag.Bool("run", false, "run server")
	flag.Parse()

	if *doUpdate {
		if *doUpdateLocation {
			update.Location()
			// locationUpdaer := update.LocationUpdater{}
			// locationUpdaer.Update()
		} else {
			update.Do()
		}
	}

	if *runApp == true {
		service.Run()
	}
}
