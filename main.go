package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ReadFromFile reads the file at filepath and returns the contents as a string.
func ReadFromFile(filepath string) string {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		log.Fatalf("open %s: %v", filepath, err)
	}

	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return string(dat)
}

// StringToLines separates a string s into an array of strings, separating at
// newlines.
func StringToLines(s string) []string {
	return strings.Split(s, "\n")
}
