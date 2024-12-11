package main

import (
	"math"
	"slices"
)

func day8() {
	type Pos struct {
		x int
		y int
	}

	puzzleMap := [][]rune{}
	antennas := map[rune][]Pos{}
	antinodes := []Pos{}

	for y, line := range readLines("input08.txt") {
		puzzleMap = append(puzzleMap, []rune{})

		for x, rune := range line {
			puzzleMap[y] = append(puzzleMap[y], rune)
			if rune != '.' {
				antennas[rune] = append(antennas[rune], Pos{x, y})
			}
		}
	}

	for _, poses := range antennas {
		for pos1Idx, pos1 := range poses {
			for pos2Idx, pos2 := range poses {
				if pos1Idx == pos2Idx {
					continue
				}

				distX := math.Abs(float64(pos1.x) - float64(pos2.x))
				distY := math.Abs(float64(pos1.y) - float64(pos2.y))

				dirX := -1
				dirY := -1

				if pos1.x > pos2.x {
					dirX = 1
				}

				if pos1.y > pos2.y {
					dirY = 1
				}

				antinodeX := pos1.x
				antinodeY := pos1.y

				for antinodeX >= 0 && antinodeY >= 0 && antinodeY < len(puzzleMap) && antinodeX < len(puzzleMap[0]) {
					if !slices.ContainsFunc(antinodes, func(pos Pos) bool {
						return pos.x == antinodeX && pos.y == antinodeY
					}) {
						antinodes = append(antinodes, Pos{antinodeX, antinodeY})
					}

					antinodeX += int(distX)*dirX
					antinodeY += int(distY)*dirY
				}
			}
		}
	}

	for y, line := range puzzleMap {
		for x, rune := range line {
			posIdx := slices.IndexFunc(antinodes, func(pos Pos) bool {
				return pos.x == x && pos.y == y
			})

			if posIdx >= 0 {
				print("#")
			} else {
				print(string(rune))
			}
		}

		println()
	}

	println(len(antinodes))
}
