package main

import (
	"os"
	"strconv"
	"strings"
)

func day11() {
	file, err := os.ReadFile("input11.txt")
	if err != nil {
		panic(err)
	}

	stones := []int{}
	for _, stonesRaw := range strings.Split(strings.TrimSpace(string(file)), " ") {
		stone, err := strconv.Atoi(stonesRaw)
		if err != nil {
			panic(err)
		}

		stones = append(stones, stone)
	}

	count := 0
	for _, stone := range stones {
		count += processStone(stone, 1)
	}

	println(count)
}

var maxBlinks = 25
var stonesCache = map[struct {
	blink int
	stone int
}]int{}

func processStone(stone, blink int) (count int) {
	if blink == maxBlinks+1 {
		return 0
	}

	cacheHit := false
	if count, cacheHit = stonesCache[struct {
		blink int
		stone int
	}{blink, stone}]; cacheHit {
		return count
	}

	if stone == 0 {
		count = processStone(1, blink+1)
	} else {
		stoneRaw := strconv.Itoa(stone)
		if len(stoneRaw)%2 == 0 {
			leftPart := stoneRaw[0 : len(stoneRaw)/2]
			rightPart := stoneRaw[len(stoneRaw)/2:]

			leftStone, err := strconv.Atoi(leftPart)
			if err != nil {
				panic(err)
			}

			rightStone, err := strconv.Atoi(rightPart)
			if err != nil {
				panic(err)
			}

			count = processStone(leftStone, blink+1) + processStone(rightStone, blink+1)
		} else {
			count = processStone(stone*2024, blink+1)
		}
	}

	stonesCache[struct {
		blink int
		stone int
	}{blink, stone}] = count

	return count
}
