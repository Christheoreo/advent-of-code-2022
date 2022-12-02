package day02

import (
	"strings"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

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
		} else if outcome == "X" { // loos3
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
