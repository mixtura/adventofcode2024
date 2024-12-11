package main

import (
	"os"
	"strings"
)

func day9() {
	file, err := os.ReadFile("input09.txt")

	if err != nil {
		panic(err)
	}

	compactDiscMap := []int{}
	moves := map[int][]int{}

	for _, rune := range strings.TrimSpace(string(file)) {
		blocksCount := int(rune - '0')
		compactDiscMap = append(compactDiscMap, blocksCount)
	}

	for fileIdx := len(compactDiscMap) - 1; fileIdx >= 0; fileIdx -= 2 {
		fileBlocksCount := compactDiscMap[fileIdx]
		for spaceIdx := 1; spaceIdx < fileIdx; spaceIdx += 2 {
			spaceBlocksCount := compactDiscMap[spaceIdx]

			if spaceBlocksCount >= fileBlocksCount {
				compactDiscMap[spaceIdx] -= fileBlocksCount
				compactDiscMap[fileIdx - 1] += fileBlocksCount

				moves[spaceIdx] = append(moves[spaceIdx], fileIdx)

				break
			}
		}
	}

	discMap := []int{}
	for idx, blocksCount := range compactDiscMap {
		element := -1

		if idx%2 == 0 {
			element = idx / 2
		} else {
			if moves, ok := moves[idx]; ok {
				for _, fileIdx := range moves {
					fileBlocksCount := compactDiscMap[fileIdx]
					compactDiscMap[fileIdx] -= fileBlocksCount

					for fileBlocksCount > 0 {
						discMap = append(discMap, fileIdx / 2)
						fileBlocksCount--
					}
				}
			}
		}

		for blocksCount > 0 {
			discMap = append(discMap, element)
			blocksCount--
		}
	}

	checksum := 0

	for idx, el := range discMap {
		if el == -1 {
			print(".")
		} else {
			checksum += idx * el
			print(el)
		}
	}

	println()
	println(checksum)
}
