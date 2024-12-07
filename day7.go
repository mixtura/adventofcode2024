package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day7() {
	lines := readLines("input7.txt")
	total := 0

	for _, line := range lines {
		sumRaw, numsRaw, ok := strings.Cut(line, ":")
		if !ok {
			panic("invalid input")
		}

		testValue, err := strconv.Atoi(sumRaw)
		if err != nil {
			panic(err)
		}

		nums := []int{}
		for _, numRaw := range strings.Split(strings.TrimSpace(numsRaw), " ") {
			num, err := strconv.Atoi(numRaw)

			if err != nil {
				panic(err)
			}

			nums = append(nums, num)
		}

		if getEquationsCount(testValue, nums[0], nums[1:], 0) ||
			getEquationsCount(testValue, nums[0], nums[1:], 1) ||
			getEquationsCount(testValue, nums[0], nums[1:], 2) {

			println(testValue)
			total += testValue
		}
	}

	println(total)
}

func getEquationsCount(testValue int, acc int, nums []int, op int) bool {
	if op == 0 {
		acc += nums[0]
	}

	if op == 1 {
		acc *= nums[0]
	}

	if op == 2 {
		var err error
		acc, err = strconv.Atoi(fmt.Sprintf("%d%d", acc, nums[0]))

		if err != nil {
			panic(err)
		}
	}

	if acc > testValue {
		return false
	}

	if len(nums) > 1 {
		return getEquationsCount(testValue, acc, nums[1:], 0) ||
			getEquationsCount(testValue, acc, nums[1:], 1) ||
			getEquationsCount(testValue, acc, nums[1:], 2)
	}

	if testValue == acc {
		return true
	}

	return false
}
