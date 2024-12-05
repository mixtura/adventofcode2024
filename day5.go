package main

import (
	"os"
	"strconv"
	"strings"
)

func day5() {
	file, err := os.ReadFile("input5.txt")
	if err != nil {
		panic(err)
	}

	text := string(file)
	rulesPart, updatesPart, ok := strings.Cut(text, "\n\n")
	if !ok {
		panic("invalid input")
	}

	compressedRules := map[int][]int{}

	for _, ruleLine := range strings.Split(rulesPart, "\n") {
		rule := strings.SplitN(ruleLine, "|", 2)
		ruleNum1, err := strconv.Atoi(rule[0])
		if err != nil {
			panic(err)
		}

		ruleNum2, err := strconv.Atoi(rule[1])
		if err != nil {
			panic(err)
		}

		compressedRules[ruleNum1] = append(compressedRules[ruleNum1], ruleNum2)
	}

	updates := [][]int{}
	for _, updateLine := range strings.Split(strings.TrimSpace(updatesPart), "\n") {
		update := []int{}
		for _, numRaw := range strings.Split(updateLine, ",") {
			num, err := strconv.Atoi(numRaw)
			if (err != nil) {
				panic(err)
			}

			update = append(update, num)
		}

		updates = append(updates, update)
	}

	total := 0
	for _, update := range updates {
		if checkUpdate(update, compressedRules) {
			total += update[len(update) / 2]
		}
	}

	println(total)
}

func checkUpdate(update []int, compressedRules map[int][]int) bool {
	for numIdx, num := range update {
		if numsAfter, ok := compressedRules[num]; ok {
			for idxToCheck := numIdx - 1; idxToCheck >= 0; idxToCheck-- {
				for _, numAfter := range numsAfter {
					if update[idxToCheck] == numAfter {
						return false
					}
				}
			}	
		}

	}

	return true
}
