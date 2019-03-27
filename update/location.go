package update

type location struct {
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
