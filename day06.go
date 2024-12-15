package main

import (
	"strings"
	"time"
)

var puzzleMap = [][]rune{}
var puzzlePath = [][]int{}
var dirs = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func day6() {

	guardDirRunes := "^>v<"
	pathDirRunes := []rune("|-|-")

	guardPosX := 0
	guardPosY := 0
	guardDirIdx := 0

	for lineIdx, line := range readLines("input06.txt") {
		puzzleMap = append(puzzleMap, []rune{})
		puzzlePath = append(puzzlePath, []int{})
		for runeIdx, rune := range line {
			puzzleMap[lineIdx] = append(puzzleMap[lineIdx], rune)
			puzzlePath[lineIdx] = append(puzzlePath[lineIdx], -1)

			dirRuneIdx := strings.Index(guardDirRunes, string(rune))

			if dirRuneIdx >= 0 {
				guardDirIdx = dirRuneIdx
				guardPosX = runeIdx
				guardPosY = lineIdx

				puzzleMap[lineIdx][runeIdx] = pathDirRunes[guardDirIdx]
				puzzlePath[lineIdx][runeIdx] = guardDirIdx
			}
		}
	}

	posX := guardPosX
	posY := guardPosY

	drawMap := func() {
		time.Sleep(100000000)
		clearConsole()
		for _, line := range puzzleMap {
			for _, rune := range line {
				print(string(rune))
			}
			println()
		}
	}

	tryFindLoop := func(startPosX, startPosY int, dirIdx int) bool {
		posX := startPosX
		posY := startPosY

		for {
			nextPosX := posX + dirs[dirIdx%4][0]
			nextPosY := posY + dirs[dirIdx%4][1]

			if nextPosX < 0 || nextPosY < 0 || nextPosY >= len(puzzleMap) || nextPosX >= len(puzzleMap[0]) {
				break
			}

			rune := puzzleMap[nextPosY][nextPosX]

			if rune == '#' {
				dirIdx++
				continue
			} else {
				posX = nextPosX
				posY = nextPosY
			}

			if puzzleMap[posY][posX] == '+' || puzzlePath[posY][posX] == dirIdx%4 {
				return true
			}

			puzzlePath[posY][posX] = dirIdx % 4
		}

		return false
	}

	obstacles := 0
	steps := 0

	for {
		dir := dirs[guardDirIdx%4]
		nextPosX := posX + dir[0]
		nextPosY := posY + dir[1]

		if nextPosX < 0 || nextPosY < 0 || nextPosY >= len(puzzleMap) || nextPosX >= len(puzzleMap[0]) {
			break
		}

		rune := puzzleMap[nextPosY][nextPosX]

		if rune == '#' {
			guardDirIdx++
		} else {
			if puzzleMap[nextPosY][nextPosX] != 'O' {
				puzzleMap[nextPosY][nextPosX] = pathDirRunes[guardDirIdx%4]
			}

			posX = nextPosX
			posY = nextPosY

			nextObstacleX := posX + dir[0]
			nextObstacleY := posY + dir[1]

			if nextObstacleX >= 0 && nextObstacleY >= 0 && nextObstacleY < len(puzzleMap) && nextObstacleX < len(puzzleMap[0]) &&
				puzzleMap[nextObstacleY][nextObstacleX] != '#' &&
				tryFindLoop(nextPosX, nextPosY, guardDirIdx+1) {
				obstacles++
				puzzleMap[nextObstacleY][nextObstacleX] = 'O'
			}

			puzzlePath[nextPosY][nextPosX] = guardDirIdx % 4

			steps++
		}

		drawMap()
	}

	drawMap()
	println(steps)
	println(obstacles)
}

// func day6() {
// 	puzzleMap := [][]rune{}
// 	puzzlePath := [][]int{}
//
// 	guardDirRunes := "^>v<"
// 	pathDirRunes := []rune("|-|-")
// 	dirs := [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
//
// 	guardPosX := 0
// 	guardPosY := 0
// 	guardDirIdx := 0
//
// 	for lineIdx, line := range readLines("input6.txt") {
// 		puzzleMap = append(puzzleMap, []rune{})
// 		puzzlePath = append(puzzlePath, []int{})
// 		for runeIdx, rune := range line {
// 			puzzleMap[lineIdx] = append(puzzleMap[lineIdx], rune)
// 			puzzlePath[lineIdx] = append(puzzlePath[lineIdx], -1)
//
// 			dirRuneIdx := strings.Index(guardDirRunes, string(rune))
//
// 			if dirRuneIdx >= 0 {
// 				guardDirIdx = dirRuneIdx
// 				guardPosX = runeIdx
// 				guardPosY = lineIdx
//
// 				puzzleMap[lineIdx][runeIdx] = pathDirRunes[guardDirIdx]
// 				puzzlePath[lineIdx][runeIdx] = guardDirIdx
// 			}
// 		}
// 	}
//
// 	obstaclesCount := 0
// 	stepsCount := 0
//
// 	drawMap := func() {
// 		time.Sleep(100000000)
// 		clearConsole()
// 		for _, line := range puzzleMap {
// 			for _, rune := range line {
// 				print(string(rune))
// 			}
// 			println()
// 		}
// 	}
//
// 	for {
// 		dir := dirs[guardDirIdx%4]
// 		nextPosX := guardPosX + dir[0]
// 		nextPosY := guardPosY + dir[1]
//
// 		traceBack := func() {
// 			posX := guardPosX
// 			posY := guardPosY
// 			opositeDir := dirs[(guardDirIdx+2)%4]
//
// 			for {
// 				posX += opositeDir[0]
// 				posY += opositeDir[1]
//
// 				if posX < 0 || posY < 0 || posY >= len(puzzleMap) ||
// 					posX >= len(puzzleMap[0]) || puzzleMap[posY][posX] == '#' {
// 					break
// 				}
//
// 				// if puzzleMap[posY][posX] != 'O' {
// 				// 	puzzleMap[posY][posX] = pathDirRunes[guardDirIdx%4]
// 				// }
//
// 				puzzlePath[posY][posX] = guardDirIdx % 4
// 			}
// 		}
//
// 		traceBack()
//
// 		if nextPosX < 0 || nextPosY < 0 || nextPosY >= len(puzzleMap) || nextPosX >= len(puzzleMap[0]) {
// 			break
// 		}
//
// 		rune := puzzleMap[nextPosY][nextPosX]
// 		if rune == '#' {
// 			guardDirIdx++
// 			if puzzleMap[guardPosY][guardPosX] != 'O' {
// 				puzzleMap[guardPosY][guardPosX] = '+'
// 			}
//
// 			traceBack()
// 			drawMap()
//
// 		} else {
// 			nextRune := puzzleMap[nextPosY][nextPosX]
// 			nextStep := puzzlePath[nextPosY][nextPosX]
//
// 			if nextStep >= 0 && nextStep == (guardDirIdx+1)%4 {
// 				obstaclesCount++
// 				puzzleMap[nextPosY+dir[1]][nextPosX+dir[0]] = 'O'
// 			} else {
// 				stepsCount++
// 			}
//
// 			if nextRune != 'O' {
// 				puzzleMap[nextPosY][nextPosX] = pathDirRunes[guardDirIdx%4]
// 			}
//
// 			puzzlePath[nextPosY][nextPosX] = guardDirIdx % 4
// 			guardPosX = nextPosX
// 			guardPosY = nextPosY
// 		}
// 	}
//
// 	drawMap()
// 	println(stepsCount)
// 	println(obstaclesCount)
// }

