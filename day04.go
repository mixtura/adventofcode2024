package main

var foundLetters [][]rune

func drawFoundLetters() {
	for _, r1 := range foundLetters {
		println()
		for _, r2 := range r1 {
			print(string(r2))
		}
	}
}

func resetFoundLetters() {
	for _, line := range foundLetters {
		for idx := range line {
			line[idx] = ' '
		}
	}
}

func day4() {
	letters := [][]rune{}

	for lineIdx, line := range readLines("input04.txt") {
		letters = append(letters, []rune{})
		foundLetters = append(foundLetters, []rune{})

		println()
		for _, rune := range line {
			print(string(rune))

			letters[lineIdx] = append(letters[lineIdx], rune)
			foundLetters[lineIdx] = append(foundLetters[lineIdx], ' ')
		}
	}

	dirs := [][]int{
		{-1, 1},
		{-1, -1},
		{1, 1},
		{1, -1},

		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	total := 0
	for xIdx := 0; xIdx < len(letters); xIdx++ {
		for yIdx := 0; yIdx < len(letters[xIdx]); yIdx++ {
			for _, dir := range dirs {
				if checkWord("XMAS", letters, xIdx, yIdx, dir[0], dir[1]) {
					total += 1
				}
			}
		}
	}

	println()
	drawFoundLetters()
	resetFoundLetters()
	println(total)

	total = 0
	for xIdx := 0; xIdx < len(letters); xIdx++ {
		for yIdx := 0; yIdx < len(letters[xIdx]); yIdx++ {
			if checkXMas(letters, xIdx, yIdx) {
				total += 1
			}
		}
	}

	println()
	drawFoundLetters()
	println(total)
}

func checkXMas(letters [][]rune, x, y int) bool {
	firstMasX := x
	firstMasY := y

	secondMasX := x + 2
	secondMasY := y

	return (checkWord("MAS", letters, firstMasX, firstMasY, 1, -1) ||
		checkWord("SAM", letters, firstMasX, firstMasY, 1, -1)) &&
		(checkWord("MAS", letters, secondMasX, secondMasY, -1, -1) ||
			checkWord("SAM", letters, secondMasX, secondMasY, -1, -1))
}

func checkWord(word string, letters [][]rune, x, y int, xDir, yDir int) bool {
	found := []int{}

	for _, r := range word {
		if x >= len(letters) || x < 0 || y >= len(letters[x]) || y < 0 {
			return false
		}

		if letters[x][y] != r {
			return false
		}

		found = append(found, x, y)

		x += xDir
		y += yDir
	}

	for idx := 0; idx < len(found); idx += 2 {
		foundLetters[found[idx]][found[idx+1]] = letters[found[idx]][found[idx+1]]
	}

	return true
}

