package service

import (
	"log"
	"net/http"
	"os"
	"time"
)

type Service struct {
	Port string
}

type FileInfo struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}

func (s *Service) Init() {
	http.HandleFunc("/api/sites", ApiSites())

	// http.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
	// 	output := map[string]interface{}{
	// 		"test": "test",
	// 	}

	// 	// 拿到最新的一隻
	// 	entries, err := ioutil.ReadDir("./json")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	list := []FileInfo{}

	// 	for _, entry := range entries {
	// 		f := FileInfo{
	// 			Name:    entry.Name(),
	// 			Size:    entry.Size(),
	// 			Mode:    entry.Mode(),
	// 			ModTime: entry.ModTime(),
	// 			IsDir:   entry.IsDir(),
	// 		}

	// 		list = append(list, f)
	// 	}

	// 	output["files"] = &list

	// 	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	// 	w.Header().Set("Server", "A Go Web Server")
	// 	w.WriteHeader(http.StatusOK)

	// 	outputByte, err := json.Marshal(output)
	// 	w.Write(outputByte)
	// })

	log.Fatal(http.ListenAndServe(s.Port, nil))
}

func (s *Service) SetPort(port string) *Service {
	s.Port = port
	return s
}

func Run() {
	s := new(Service)
	s.SetPort(":6543")
	s.Init()
}
