package main

import (
	"math"
	"regexp"
	"strconv"
)

func day14() {
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	height := 103
	width := 101
	time := 1
	minSafety := math.MaxInt

	for {
		firstQuarter := 0
		secondQuarter := 0
		thirdQuarter := 0
		forthQuarter := 0

		for _, line := range readLines("input14.txt") {
			matches := re.FindStringSubmatch(line)
		
			if len(matches) == 5 {
				px, _ := strconv.Atoi(matches[1])
				py, _ := strconv.Atoi(matches[2])
				vx, _ := strconv.Atoi(matches[3])
				vy, _ := strconv.Atoi(matches[4])

				finalPx := (px + time * vx) % width
				finalPy := (py + time * vy) % height

				if finalPx < 0 {
					finalPx = width + finalPx
				}

				if finalPy < 0 {
					finalPy = height + finalPy
				}

				if finalPx == width / 2 || finalPy == height / 2 {
					continue
				}

				if finalPx < width / 2 {
					if finalPy < height / 2 {
						firstQuarter++
					} else {
						forthQuarter++
					}
				} else {
					if finalPy < height / 2 {
						secondQuarter++
					} else {
						thirdQuarter++
					}
				}
			}
		}

		safety := firstQuarter * secondQuarter * thirdQuarter * forthQuarter

		if minSafety > safety {
			minSafety = safety
			println(minSafety)
			println(time)
			println()
		}

		time++
	}
}
