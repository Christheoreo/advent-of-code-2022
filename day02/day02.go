package day02

import (
	"strings"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

// #4 is by far the fastest

func SolveFirst(filename string) int {
	// a = rock, b = paper, c = scissors. x = rock, y = paper, z = scissors
	data, _ := filereader.ReadFileToStringArray(filename)
	runningTotal := 0
	for _, v := range data {
		parts := strings.Split(v, " ")
		points := 0

		them := parts[0]
		you := parts[1]

		if you == "X" {
			points += 1
			if them == "A" {
				points += 3
			} else if them == "C" {
				points += 6
			}
		} else if you == "Y" {
			points += 2
			if them == "B" {
				points += 3
			} else if them == "A" {
				points += 6
			}
		} else {
			points += 3
			if them == "C" {
				points += 3
			} else if them == "B" {
				points += 6
			}
		}

		runningTotal += points

	}

	return runningTotal
}
func SolveSecond(filename string) int {
	// a = rock, b = paper, c = scissors. x = loose, y = draw, z = win
	data, _ := filereader.ReadFileToStringArray(filename)
	runningTotal := 0
	for _, v := range data {
		parts := strings.Split(v, " ")
		points := 0

		them := parts[0]
		outcome := parts[1]

		// need to Draw
		if outcome == "Y" {
			points += 3
			if them == "A" {
				points += 1
			} else if them == "B" {
				points += 2
			} else if them == "C" {
				points += 3
			}
		} else if outcome == "X" { // loose
			points += 0
			if them == "A" {
				points += 3
			} else if them == "B" {
				points += 1
			} else if them == "C" {
				points += 2
			}
		} else { //win
			points += 6
			if them == "A" {
				points += 2
			} else if them == "B" {
				points += 3
			} else if them == "C" {
				points += 1
			}
		}
		runningTotal += points
	}
	return runningTotal
}

// This looks nicer
func SolveSecondAttempt2(filename string) int {
	// a = rock, b = paper, c = scissors. x = lose, y = draw, z = win
	data, _ := filereader.ReadFileToStringArray(filename)
	runningTotal := 0
	mapper := map[string]map[string]int{
		"Y": {
			"A": 1,
			"B": 2,
			"C": 3,
		},
		"X": {
			"A": 3,
			"B": 1,
			"C": 2,
		},
		"Z": {
			"A": 2,
			"B": 3,
			"C": 1,
		},
	}
	for _, v := range data {
		parts := strings.Split(v, " ")
		points := 0

		them := parts[0]
		outcome := parts[1]

		if outcome == "Y" { // Draw
			points += 3
		} else if outcome == "Z" { //win
			points += 6
		}
		points += mapper[outcome][them]
		runningTotal += points
	}
	return runningTotal
}

func SolveSecondAttempt3(filename string) int {
	// a = rock, b = paper, c = scissors. x = loose, y = draw, z = win
	data, _ := filereader.ReadFileToStringArray(filename)
	runningTotal := 0

	all := map[string]map[string]int{
		"X": {
			"A": 3,
			"B": 1,
			"C": 2,
		},
		"Y": { //added extra 3
			"A": 4,
			"B": 5,
			"C": 6,
		},
		"Z": { // added extra 6
			"A": 8,
			"B": 9,
			"C": 7,
		},
	}

	for _, v := range data {
		parts := strings.Split(v, " ")
		points := 0

		them := parts[0]
		outcome := parts[1]

		points += all[outcome][them]
		runningTotal += points
	}
	return runningTotal
}

// This is the fastest
func SolveSecondAttempt4(filename string) int {
	// a = rock, b = paper, c = scissors. x = loose, y = draw, z = win
	data, _ := filereader.ReadFileToStringArray(filename)
	runningTotal := 0
	for _, v := range data {
		if v == "A X" {
			runningTotal += 3
			continue
		}
		if v == "A Y" {
			runningTotal += 4
			continue
		}

		if v == "A Z" {
			runningTotal += 8
			continue
		}

		if v == "B X" {
			runningTotal += 1
			continue
		}

		if v == "B Y" {
			runningTotal += 5
			continue
		}

		if v == "B Z" {
			runningTotal += 9
			continue
		}

		if v == "C X" {
			runningTotal += 2
			continue
		}
		if v == "C Y" {
			runningTotal += 6
			continue
		}

		if v == "C Z" {
			runningTotal += 7
			continue
		}
	}
	return runningTotal
}
