package day04

import (
	"strconv"
	"strings"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

func SolveFirst(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	pairs := 0

	for _, row := range data {
		sections := strings.Split(row, ",")

		a, b := sections[0], sections[1]

		partsA := strings.Split(a, "-")
		partsB := strings.Split(b, "-")

		a1, _ := strconv.Atoi(partsA[0])
		a2, _ := strconv.Atoi(partsA[1])

		b1, _ := strconv.Atoi(partsB[0])
		b2, _ := strconv.Atoi(partsB[1])

		if (b1 >= a1 && b2 <= a2) || (a1 >= b1 && a2 <= b2) {
			pairs++
			continue
		}

	}

	return pairs
}

func SolveSecond(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	pairs := 0

	for _, row := range data {
		sections := strings.Split(row, ",")

		a, b := sections[0], sections[1]

		partsA := strings.Split(a, "-")
		partsB := strings.Split(b, "-")

		a1, _ := strconv.Atoi(partsA[0])
		a2, _ := strconv.Atoi(partsA[1])

		b1, _ := strconv.Atoi(partsB[0])
		b2, _ := strconv.Atoi(partsB[1])

		if (a1 >= b1 && a1 <= b2) || (a2 >= b1 && a2 <= b2) || (b1 >= a1 && b1 <= a2) || (b2 >= a1 && b2 <= a2) {
			pairs++
			continue
		}

	}

	return pairs
}
