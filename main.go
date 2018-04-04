package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Config Configuration struct (will be updated as app is built)
type Config struct {
	APIKey   string `json:"api-key"`
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

// URLs Base URLs for TMDb.org
type URLs struct {
	MovieURL  string `json:"movie-url"`
	TVURL     string `json:"tv-url"`
	ImageURL  string `json:"image-url"`
	ConfigURL string `json:"config-url"`
}

// LoadConfiguration Loads configuration file (e.g. keys.json)
func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
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

func main() {

	config, _ := LoadConfiguration("keys.json")

	addr, _ := LoadURLs("URLs.json")

	url := addr.ConfigURL + "?api_key=" + config.APIKey

	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("GET", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println("------")
	fmt.Println(string(body))

}
