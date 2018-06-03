package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	. "mini-movie-database/config"
)

// URLs Base URLs for TMDb.org
type URLs struct {
	MovieURL  string `json:"movie-url"`
	TVURL     string `json:"tv-url"`
	ImageURL  string `json:"image-url"`
	ConfigURL string `json:"config-url"`
}

// LoadURLs Loads URLs from the URLs.json file
func LoadURLs(file string) (URLs, error) {
	var retVal URLs
	urlFile, err := os.Open(file)
	defer urlFile.Close()
	if err != nil {
		return retVal, err
	}
	jsonParser := json.NewDecoder(urlFile)
	jsonParser.Decode(&retVal)
	return retVal, err
}

func DumpLoad(file string) {
	movies, err := LoadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {

	config := Config

	addr, _ := LoadURLs("URLs.json")

	url := addr.ConfigURL + "?api_key=" + config.APIKey

	payload := strings.NewReader("{}")

	DumpLoad("movie_ids_04_28_2017.json")

	req, _ := http.NewRequest("GET", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println("------")
	fmt.Println(string(body))

}
