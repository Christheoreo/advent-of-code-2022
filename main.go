package main

import (
	"fmt"
	"time"

	"github.com/christheoreo/advent-of-code-2022/internal/timetrack"
	"github.com/christheoreo/advent-of-code-2022/problems"
)

var problemsMap map[int]problems.Problem = map[int]problems.Problem{
	1: problems.ProblemOne{},
}

func main() {
	fmt.Println("Advent of code 2022!")
	solveDay(1)
}

func solveDay(problemNumber int) {

	problem, ok := problemsMap[problemNumber]

	if !ok {
		panic("could not find problem to solve")
	}

	firstTime := time.Now()

	// defer TimeTrack(time.Now(), fmt.Sprintf("Problem %d", problemNumber))
	err := problem.SolveFirst()

	timetrack.TimeTrack(firstTime, fmt.Sprintf("%d - part 1", problemNumber))
	if err != nil {
		panic(err)
	}

	secondTime := time.Now()
	err = problem.SolveFirst()
	timetrack.TimeTrack(secondTime, fmt.Sprintf("%d - part 2", problemNumber))
	if err != nil {
		panic(err)
	}

}
