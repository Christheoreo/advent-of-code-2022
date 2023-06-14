package day02

import (
	"strings"
)

func SolveFirst(data string) int {
	score := 0
	for _, line := range strings.Split(data, "\n") {
		if line == "A X" {
			score += 4
		} else if line == "A Y" {
			score += 8
		} else if line == "A Z" {
			score += 3
		} else if line == "B X" {
			score += 1
		} else if line == "B Y" {
			score += 5
		} else if line == "B Z" {
			score += 9
		} else if line == "C X" {
			score += 7
		} else if line == "C Y" {
			score += 2
		} else if line == "C Z" {
			score += 6
		}
	}
	return score
}

func SolveSecond(data string) int {
	// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win.
	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	score := 0
	// A = rock, B = Paper = C = Scissors
	// 1 for rock, 2 for paper, 3 for scissors
	for _, line := range strings.Split(data, "\n") {
		if line == "A X" {
			// We need to lose against rock
			score += 3
		} else if line == "A Y" {
			score += 4
		} else if line == "A Z" {
			score += 8
		} else if line == "B X" {
			// we need to lose against paper
			score += 1
		} else if line == "B Y" {
			// We need to draw with paper
			score += 5
		} else if line == "B Z" {
			// We need to win against paper
			score += 9
		} else if line == "C X" {
			// We need to lose against scissors
			score += 2
		} else if line == "C Y" {
			// we need to draw against scissors
			score += 6
		} else if line == "C Z" {
			// We need to win against scissors
			score += 7
		}
	}
	return score
}
