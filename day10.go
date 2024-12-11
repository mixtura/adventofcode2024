package main

import "slices"

var trailEnds []int

func day10() {
	puzzleMap := [][]int{}

	for lineIdx, line := range readLines("input10.txt") {
		puzzleMap = append(puzzleMap, []int{})
		for _, rune := range line {
			num := int(rune - '0')
			puzzleMap[lineIdx] = append(puzzleMap[lineIdx], num)

			print(num)
		}
		println()
	}

	totalScore := 0
	totalRating := 0
	for y, line := range puzzleMap {
		for x, h := range line {
			if h == 0 {
				trailEnds = []int{}

				score, rating := searchPath(puzzleMap, x, y, 1, 0, 0)
				totalScore += score
				totalRating += rating

				score, rating = searchPath(puzzleMap, x, y, 0, 1, 0)
				totalScore += score
				totalRating += rating

				score, rating = searchPath(puzzleMap, x, y, 0, -1, 0)
				totalScore += score
				totalRating += rating

				score, rating = searchPath(puzzleMap, x, y, -1, 0, 0)
				totalScore += score
				totalRating += rating
			}
		}
	}

	println(totalScore)
	println(totalRating)
}

func searchPath(puzzleMap [][]int, x, y int, xDir, yDir int, height int) (score int, rating int) {
	for {
		x += xDir
		y += yDir
		height++

		if y < 0 || x < 0 || y >= len(puzzleMap) || x >= len(puzzleMap[0]) || puzzleMap[y][x] != height {
			break
		}

		if height == 9 {
			rating += 1

			idx := len(puzzleMap) * x + y
			if !slices.Contains(trailEnds, idx) {
				score += 1
				trailEnds = append(trailEnds, idx)
			}

			break
		}

		// clockwise rotation:
		// -1  0
		//  0 -1
		//  1  0
		//  0  1
		if xDir != 0 {
			pathScore, pathRating := searchPath(puzzleMap, x, y, 0, -1, height)
			score += pathScore
			rating += pathRating

			pathScore, pathRating = searchPath(puzzleMap, x, y, 0, 1, height)
			score += pathScore
			rating += pathRating
		} else {
			pathScore, pathRating := searchPath(puzzleMap, x, y, 1, 0, height)
			score += pathScore
			rating += pathRating

			pathScore, pathRating = searchPath(puzzleMap, x, y, -1, 0, height)
			score += pathScore
			rating += pathRating
		}
	}

	return
}
