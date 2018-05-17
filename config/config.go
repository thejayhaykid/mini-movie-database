package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config Configuration struct (will be updated as app is built)
type Config struct {
	APIKey   string `json:"api-key"`
	Database struct {
		Server   string `json:"server"`
		Db       string `json:"database"`
		Password string `json:"password"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

// Read and parse keys.json
func (c *Config) Read() {
	configFile, err := os.Open("keys.json")
	defer configFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&c)
}
