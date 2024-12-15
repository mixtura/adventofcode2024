package main

import (
	"fmt"
)

func day12() {
	puzzleMap := readMap("input12.txt")
	traversed := make([][]bool, len(puzzleMap))

	for i := range traversed {
		traversed[i] = make([]bool, len(puzzleMap[0]))
	}

	findNextStartPoint := func() (int, int) {
		for y := range puzzleMap {
			for x := range puzzleMap[y] {
				if !traversed[y][x] {
					return x, y
				}
			}
		}

		return -1, -1
	}

	getCornersCount := func(x, y int, expectedPlant rune) int {
		dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

		count := 0
		cornerSides := 0

		for i := 0; i <= len(dirs); i++ {
			dir := dirs[(i)%len(dirs)]

			y0 := y + dir[1]
			x0 := x + dir[0]

			if y0 < 0 || x0 < 0 || y0 >= len(puzzleMap) || x0 >= len(puzzleMap[0]) || puzzleMap[y0][x0] != expectedPlant {
				cornerSides++
			} else {
				cornerSides = 0
			}

			if cornerSides >= 2 {
				count++
			}
		}

		diagonals := [][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

		for _, dir := range diagonals {
			y0 := y + dir[1]
			x0 := x + dir[0]

			y1 := y0
			x1 := x

			y2 := y
			x2 := x0

			if y0 >= 0 && x0 >= 0 && y0 < len(puzzleMap) && x0 < len(puzzleMap[0]) && puzzleMap[y0][x0] != expectedPlant {
				if puzzleMap[y1][x1] == expectedPlant && puzzleMap[y2][x2] == expectedPlant {
					count++
				}
			}
		}

		return count
	}

	var traverseMap func(x, y int, expectedPlant rune) (int, int)
	traverseMap = func(x, y int, expectedPlant rune) (int, int) {
		if x < 0 || y < 0 || x >= len(puzzleMap[0]) || y >= len(puzzleMap) || puzzleMap[y][x] != expectedPlant || traversed[y][x] {
			return 0, 0
		}

		traversed[y][x] = true

		area1, corners1 := traverseMap(x+1, y, expectedPlant)
		area2, corners2 := traverseMap(x, y+1, expectedPlant)
		area3, corners3 := traverseMap(x-1, y, expectedPlant)
		area4, corners4 := traverseMap(x, y-1, expectedPlant)

		corners := getCornersCount(x, y, expectedPlant)

		return area1 + area2 + area3 + area4 + 1, corners1 + corners2 + corners3 + corners4 + corners
	}

	total := 0
	for {
		x, y := findNextStartPoint()

		if x < 0 {
			break
		}

		regionPlant := puzzleMap[y][x]
		area, sides := traverseMap(x, y, regionPlant)

		fmt.Printf("%s: area %d, sides %d", string(regionPlant), area, sides)
		println()

		total += area * sides
	}

	println(total)
}
