package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type tDriver struct {
	Name     string   `json:"Name"`
	Team     string   `json:"Team"`
	Nation   string   `json:"Nation"`
	GUID     string   `json:"Guid"`
	GUIDList []string `json:"GuidsLists"`
}

type tResults struct {
	Name     string `json:"DriverName"`
	GUID     string `json:"DriverGuid"`
	CarID    int    `json:"CarId"`
	CarMod   string `json:"CarModel"`
	BLap     int    `json:"BestLap"` // In 1/100s
	ToTime   int    `json:"TotalTime"`
	Ballast  int    `json:"BallastKG"`
	Restrict int    `json:"Restrictor"`
}

type tCars struct {
	CarID    int     `json:"CarId"`
	Driver   tDriver `json:"Driver"`
	Model    string  `json:"Model"`
	Skin     string  `json:"Skin"`
	Ballast  int     `json:"BallastKG"`
	Restrict int     `json:"Restrictor"`
}

type session struct {
	TrackName string     `json:"TrackName"`
	TrackConf string     `json:"TrackConfig"`
	RaceType  string     `json:"Type"`
	Duration  int        `json:"DurationSecs"`
	Laps      int        `json:"RaceLaps"`
	Cars      []tCars    `json:"Cars"`
	Results   []tResults `json:"Result"`
}

func getResults(w http.ResponseWriter, r *http.Request) {
	dirname := "./results"

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for i := range files {
		if i >= 1 {
			break
		}
		//fmt.Println(files[len(files)-1-i].Name())

		filename := dirname + "/" + files[len(files)-1-i].Name()
		fmt.Println(filename)
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
		js, err := json.Marshal(string(content))

		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func getGoldResults(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./dist")))
	http.HandleFunc("/results", getResults)
	http.ListenAndServe(":8080", nil)
}
