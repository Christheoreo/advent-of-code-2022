package problems

import "errors"

type ProblemOne struct{}

func (p ProblemOne) Solve() error {
	err := p.SolveFirst()

	if err != nil {
		return err
	}

	return p.SolveSecond()
}
func (p ProblemOne) SolveFirst() error {
	return errors.New("not done yet")
}
func (p ProblemOne) SolveSecond() error {
	return errors.New("not done yet")
}
