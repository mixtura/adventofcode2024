package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func day13() {
	re := regexp.MustCompile(`Button A: X\+(-?\d+), Y\+(-?\d+)\s*Button B: X\+(-?\d+), Y\+(-?\d+)\s*Prize: X=(-?\d+), Y=(-?\d+)`)
	puzzle, err := os.ReadFile("input13.txt")

	if err != nil {
		panic(err)
	}

	var total int64
	for _, config := range strings.Split(string(puzzle), "\n\n") {
		matches := re.FindStringSubmatch(config)

		if len(matches) == 7 {
			ax, _ := strconv.Atoi(matches[1])
			ay, _ := strconv.Atoi(matches[2])
			bx, _ := strconv.Atoi(matches[3])
			by, _ := strconv.Atoi(matches[4])
			px, _ := strconv.Atoi(matches[5])
			py, _ := strconv.Atoi(matches[6])

			px += 10000000000000
			py += 10000000000000

			fmt.Printf("Button A: X=%d, Y=%d\n", ax, ay)
			fmt.Printf("Button B: X=%d, Y=%d\n", bx, by)
			fmt.Printf("Prize: X=%d, Y=%d\n", px, py)

			anum, bnum, cost, err := solveMachine(int64(ax), int64(bx), int64(px), int64(ay), int64(by), int64(py))

			if err != nil {
				fmt.Println("Error:", err)
				continue
			} else {
				total += cost
			}

			print(anum)
			print(" ")
			print(bnum)
			println()
		} else {
			panic("No match found")
		}
	}

	println()
	println(total)
}

func solveMachine(ax, bx, px, ay, by, py int64) (int64, int64, int64, error) {
	// Solve the first equation using the Extended Euclidean Algorithm
	gcdX, x0, y0 := extendedGCD(ax, bx)
	if px%gcdX != 0 {
		return 0, 0, 0, fmt.Errorf("no solution for the first equation")
	}

	// Scale the solution for the first equation
	x0 *= px / gcdX
	y0 *= px / gcdX

	// Parameterize the solution
	kx := bx / gcdX
	ky := -ax / gcdX

	// Compute bounds for k
	kMin := int64(math.Ceil(float64(-x0) / float64(kx)))
	kMax := int64(math.Floor(float64(y0) / float64(-ky)))

	if kMin > kMax {
		return 0, 0, 0, fmt.Errorf("no valid solution exists")
	}

	// Minimize cost
	delta := int64(3*kx - ky)
	var k int64
	if delta > 0 {
		k = kMin
	} else {
		k = kMax
	}

	// Compute anum and bnum
	anum := x0 + k*kx
	bnum := y0 + k*ky

	// Validate the second equation
	if ay*anum+by*bnum != py {
		return 0, 0, 0, fmt.Errorf("solution does not satisfy the second equation")
	}

	// Compute the cost
	cost := 3*anum + bnum

	return anum, bnum, cost, nil
}

func extendedGCD(a, b int64) (int64, int64, int64) {
	if b == 0 {
		return a, 1, 0
	}

	gcd, x1, y1 := extendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}
