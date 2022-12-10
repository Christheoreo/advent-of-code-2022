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
	total := 0
	x := 1
	cycle := 1
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
	}
	return total
}

func SolveSecond(filename string) [6]string {
	data, _ := filereader.ReadFileToStringArray(filename)
	x := 1 // middle spriate pos
	cycle := 1
	rows := [6]string{}
	rowIndex := 0
	colIndex := 0
	for _, line := range data {
		// Start of the cycle
		num := 0
		iterations := 1
		if line != "noop" {
			iterations = 2
			num, _ = strconv.Atoi(strings.Split(line, " ")[1])
		}
		for i := 0; i < iterations; i++ {
			// draw the pixel
			char := "."
			if colIndex == x || colIndex == x-1 || colIndex == x+1 {
				char = "#"
			}
			rows[rowIndex] += char
			cycle++
			colIndex++

			if colIndex > 39 {
				colIndex = 0
				rowIndex++
			}
		}
		x += num
	}
	return rows
}
