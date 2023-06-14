package day01

import (
	"strconv"
	"strings"
)

func SolveFirst(data string) int {
	largest := 0
	runningTotal := 0

	for _, v := range strings.Split(data, "\n") {
		if v == "" {
			if runningTotal > largest {
				largest = runningTotal
			}
			runningTotal = 0
			continue
		}
		val, _ := strconv.Atoi(v)
		runningTotal += val
	}
	return largest
}
func SolveSecond(data string) int {
	l := [3]int{0, 0, 0}

	runningTotal := 0

	for _, v := range strings.Split(data, "\n") {
		if v == "" {
			checkRunningTotal(runningTotal, &l)
			runningTotal = 0
			continue
		}
		val, _ := strconv.Atoi(v)
		runningTotal += val
	}
	checkRunningTotal(runningTotal, &l)
	return l[0] + l[1] + l[2]
}

func checkRunningTotal(runningTotal int, l *[3]int) {
	if runningTotal > l[0] {
		l[2] = l[1]
		l[1] = l[0]
		l[0] = runningTotal
	} else if runningTotal > l[1] {
		l[2] = l[1]
		l[1] = runningTotal
	} else if runningTotal > l[2] {
		l[2] = runningTotal
	}
}
