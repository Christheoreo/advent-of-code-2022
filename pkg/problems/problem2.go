package problems

import (
	"fmt"

	"github.com/christheoreo/advent-of-code-2022/internal/problems/filereader"
)

type ProblemTwo struct{}

func (p ProblemTwo) Solve() error {
	err := p.SolveFirst()

	if err != nil {
		return err
	}

	return p.SolveSecond()
}
func (p ProblemTwo) SolveFirst() error {
	data, _ := filereader.ReadFileToStringArray("p2e.txt")
	fmt.Println(data)
	return nil
}
func (p ProblemTwo) SolveSecond() error {
	data, _ := filereader.ReadFileToStringArray("p2e.txt")
	fmt.Println(data)
	return nil
}
