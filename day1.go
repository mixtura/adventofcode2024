package main

import (
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func day1() {
	lines := readLines("input1.txt")
	leftUniqueList := map[int]bool{}
	rightCountsList := map[int]int{}

	for _, line := range lines {
		num1Raw, num2Raw, ok := strings.Cut(line, "   ")

		if !ok {
			println(line)
			continue
		}

		num1, err := strconv.Atoi(num1Raw)
		if err != nil {
			panic(err)
		}

		num2, err := strconv.Atoi(num2Raw)
		if err != nil {
			panic(err)
		}

		leftUniqueList[num1] = true
		if count, ok := rightCountsList[num2]; ok {
			rightCountsList[num2] = count + 1
		} else {
			rightCountsList[num2] = 1
		}
	}

	total := 0
	leftUniqueNums := maps.Keys(leftUniqueList)
	for _, n := range leftUniqueNums {
		total += rightCountsList[n] * n
	}

	println(total)
}
