package problems

import (
	"fmt"
	"strconv"

	"github.com/christheoreo/advent-of-code-2022/internal/problems/filereader"
)

type ProblemOne struct{}

func (p ProblemOne) Solve(filename string) error {
	err := p.SolveFirst(filename)

	if err != nil {
		return err
	}

	return p.SolveSecond(filename)
}
func (p ProblemOne) SolveFirst(filename string) error {
	data, _ := filereader.ReadFileToStringArray(filename)
	largest := 0
	runningTotal := 0

	for _, v := range data {
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

	fmt.Println(largest)
	return nil
}
func (p ProblemOne) SolveSecond(filename string) error {
	data, _ := filereader.ReadFileToStringArray(filename)
	l := [3]int{0, 0, 0}

	runningTotal := 0

	for _, v := range data {
		if v == "" {
			checkRunningTotal(runningTotal, &l)
			runningTotal = 0
			continue
		}
		val, _ := strconv.Atoi(v)
		runningTotal += val
	}
	checkRunningTotal(runningTotal, &l)
	fmt.Println(l[0] + l[1] + l[2])
	return nil
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
