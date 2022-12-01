package problems

import (
	"fmt"

	"github.com/christheoreo/advent-of-code-2022/internal/problems/filereader"
)

type ProblemTwo struct{}

func (p ProblemTwo) Solve(filename string) error {
	err := p.SolveFirst(filename)

	if err != nil {
		return err
	}

	return p.SolveSecond(filename)
}
func (p ProblemTwo) SolveFirst(filename string) error {
	data, _ := filereader.ReadFileToStringArray(filename)
	fmt.Println(data)
	return nil
}
func (p ProblemTwo) SolveSecond(filename string) error {
	data, _ := filereader.ReadFileToStringArray(filename)
	fmt.Println(data)
	return nil
}
