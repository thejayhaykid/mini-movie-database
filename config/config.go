package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Loading information from the keys.json file
type Config struct {
	TMDBKey  string
	Server   string
	Database string
}

// Read and parse keys.json
func (c *Config) Read() {
	if _, err := toml.DecodeFile("keys.json", &c); err != nil {
		log.Fatal(err)
	}
}
