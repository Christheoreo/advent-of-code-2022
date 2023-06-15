package day04

import (
	"strconv"
	"strings"
)

type CriteriaCheckFunc func(a1, a2, b1, b2 int) bool

func First(a1, a2, b1, b2 int) bool {
	return b1 >= a1 && b1 <= a2 && b2 >= a1 && b2 <= a2 ||
		a1 >= b1 && a1 <= b2 && a2 >= b1 && a2 <= b2
}

func Second(a1, a2, b1, b2 int) bool {
	return b1 >= a1 && b1 <= a2 ||
		b2 >= a1 && b2 <= a2 ||
		a1 >= b1 && a1 <= b2 ||
		a2 >= b1 && a2 <= b2
}

func SolveFirst(data string) int {
	return mainLogic(data, First)
}

func mainLogic(data string, f CriteriaCheckFunc) int {
	answer := 0
	assignmentPairsS := strings.Split(data, "\n")
	for _, pairS := range assignmentPairsS {

		var a1, a2, b1, b2 int
		parts := strings.Split(pairS, ",")

		for i, p := range parts {
			bytes := strings.Split(p, "-")
			if i == 0 {
				a1, _ = strconv.Atoi(bytes[0])
				a2, _ = strconv.Atoi(bytes[1])
				continue
			}
			b1, _ = strconv.Atoi(bytes[0])
			b2, _ = strconv.Atoi(bytes[1])
		}

		if f(a1, a2, b1, b2) {
			answer++
		}
	}
	return answer
}

func SolveSecond(data string) int {
	return mainLogic(data, Second)
}
