package problems

import (
	"errors"
	"fmt"

	"github.com/christheoreo/advent-of-code-2022/internal/problems/filereader"
)

type ProblemX struct{}

func (p ProblemX) Solve(filename string) error {
	err := p.SolveFirst(filename)

	if err != nil {
		return err
	}

	return p.SolveSecond(filename)
}
func (p ProblemX) SolveFirst(filename string) error {
	data, _ := filereader.ReadFileToStringArray(filename)
	fmt.Println(data)
	return errors.New("not done yet")
}
func (p ProblemX) SolveSecond(filename string) error {
	data, _ := filereader.ReadFileToStringArray(filename)
	fmt.Println(data)
	return errors.New("not done yet")
}
