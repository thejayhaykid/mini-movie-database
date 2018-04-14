package main

import (
	"bufio"
	"encoding/json"
	"os"
	"github.com/golang-collections/collections/stack"
)

// DBMovie is the structure for movies in the MovieDB daily dump
type DBMovie struct {
	adult bool    `json:"adult"`
	id    int64   `json:"id"`
	title string  `json:"original_title"`
	pop   float64 `json:"popularity"`
	video bool    `json:"video"`
}

// LoadFile Loads the movie from the JSON file to the struct, returns stack of movies
func LoadFile(file string) (Stack, error) {
	var movies Stack
	movieFile, err := os.Open(file)
	defer movieFile.Close()
	if err != nil {
		return movies, err
	}

	fileScanner := bufio.NewScanner(movieFile)

	for fileScanner.Scan() {
		var tempMovie DBMovie
		movieParser := json.NewDecoder(fileScanner.Text())
		movieParser.Decode(&tempMovie)
		movies.Push(tempMovie)
	}

	return movies
}
