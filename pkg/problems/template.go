package problems

import "errors"

type ProblemX struct{}

func (p ProblemX) Solve() error {
	err := p.SolveFirst()

	if err != nil {
		return err
	}

	return p.SolveSecond()
}
func (p ProblemX) SolveFirst() error {
	return errors.New("not done yet")
}
func (p ProblemX) SolveSecond() error {
	return errors.New("not done yet")
}
