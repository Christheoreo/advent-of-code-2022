package problems

import "testing"

func TestProblemXPartOne(t *testing.T) {
	p := ProblemOne{}

	err := p.SolveFirst()

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblemXPartTwo(t *testing.T) {
	p := ProblemOne{}

	err := p.SolveSecond()

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblemX(t *testing.T) {
	p := ProblemOne{}

	err := p.Solve()

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}
