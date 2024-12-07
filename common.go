package main

import (
	"os"
	"strings"
)

func readLines(path string) []string {
	file, err := os.ReadFile(path)
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	if err != nil {
		panic(err)
	}

	return lines
}

