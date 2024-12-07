package main

import (
	"slices"
	"strconv"
	"strings"
)

func day2() {
	reports := [][]int{}

	for _, line := range readLines("input2.txt") {
		if len(line) == 0 {
			continue
		}

		levelsRaw := strings.Split(line, " ")
		report := []int{}

		for _, levelRaw := range levelsRaw {
			level, err := strconv.Atoi(levelRaw)

			if err != nil {
				panic(err)
			}

			report = append(report, level)
		}

		reports = append(reports, report)
	}

	safeReportsCount := 0

	for _, report := range reports {
		tolarence := 1
		if checkReport(slices.Clone(report), tolarence) {
			safeReportsCount++
		}
	}

	println(safeReportsCount)
}

func checkReport(report []int, tolerance int) bool {
	increases := 0
	decreases := 0

	for lvlIdx := 1; lvlIdx < len(report); lvlIdx++ {
		if report[lvlIdx-1] < report[lvlIdx] {
			increases++
		} else {
			decreases++
		}
	}

	increasing := increases > decreases

	for lvlIdx := 1; lvlIdx < len(report); lvlIdx++ {
		var diff int

		level := report[lvlIdx]
		prevLevel := report[lvlIdx-1]

		if prevLevel > level {
			if increasing {
				if tolerance != 0 {
					return checkReportOptions(report, lvlIdx)
				} else {
					return false
				}
			}

			diff = prevLevel - level
		}

		if prevLevel < level {
			if !increasing {
				if tolerance != 0 {
					return checkReportOptions(report, lvlIdx)
				} else {
					return false
				}
			}

			diff = level - prevLevel
		}

		if diff == 0 || diff > 3 {
			if tolerance != 0 {
				return checkReportOptions(report, lvlIdx)
			} else {
				return false
			}
		}
	}

	return true
}

func checkReportOptions(report []int, lvlIdx int) bool {
	report1 := slices.Clone(report)
	report1 = append(report1[0:lvlIdx], report1[lvlIdx+1:]...)

	report2 := slices.Clone(report)
	report2 = append(report2[0:lvlIdx-1], report2[lvlIdx:]...)

	return checkReport(report1, 0) || checkReport(report2, 0)
}

