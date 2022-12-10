package day10

import (
	"strconv"
	"strings"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

type Instruction struct {
	Direction string
	Count     int
}

func SolveFirst(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	// cyclesToAdd := make([]int, 0)
	total := 0
	x := 1
	cycle := 1
	// queue := [3][]int{make([]int, 0), make([]int, 0), make([]int, 0)}
	for _, line := range data {
		// Start of the cycle
		num := 0
		iterations := 1
		if line != "noop" {
			iterations = 2
			num, _ = strconv.Atoi(strings.Split(line, " ")[1])
		}
		for i := 0; i < iterations; i++ {
			if cycle == 20 {
				total += (20 * x)
			}
			if cycle == 60 {
				total += (60 * x)
			}
			if cycle == 100 {
				total += (100 * x)
			}
			if cycle == 140 {
				total += (140 * x)
			}
			if cycle == 180 {
				total += (180 * x)
			}
			if cycle == 220 {
				total += (220 * x)
			}
			cycle++
		}

		x += num
		// this is the during stage - this is when we need to check what cycle this is.
		// if cycle == 20 || cycle%40 == 0 {
		// 	cyclesToAdd = append(cyclesToAdd, x)
		// }

		// after the cycle has ended

		// check to see if any of the x values need adding up

		// if len(queue[0]) < 1 {
		// 	continue
		// }

		// for _, val := range queue[1] {
		// 	x += val
		// }

	}
	return total
}

func SolveSecond(filename string) int {
	// data, _ := filereader.ReadFileToStringArray(filename)
	return 0
}
