package main

import (
	"os"
	"os/exec"
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

func readMap(path string) [][]rune {
	lines := readLines(path)
	puzzleMap := [][]rune{}

	for idx, line := range lines {
		puzzleMap = append(puzzleMap, []rune{})

		for _, rune := range line {
			puzzleMap[idx] = append(puzzleMap[idx], rune)
		}
	}

	return puzzleMap
}

func clearConsole() {
	cmd := exec.Command("clear") // Command to clear the terminal
	cmd.Stdout = os.Stdout
	cmd.Run()
}
