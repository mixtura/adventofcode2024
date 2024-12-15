package main

import (
	"os"
	"strconv"
)

func main() {
	dayNum, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic(err)
	}

	switch(dayNum) {
		case 1: day1()
		case 2: day2()
		case 3: day3()
		case 4: day4()
		case 5: day5()
		case 6: day6()
		case 7: day7()
		case 8: day8()
		case 9: day9()
		case 10: day10()
		case 11: day11()
		case 12: day12()
		case 13: day13()
		case 14: day14()
	}
}
