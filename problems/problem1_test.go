package problems

import "testing"

func TestProblemOnePartOne(t *testing.T) {
	p := ProblemOne{}

	err := p.SolveFirst()

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblemOnePartTwo(t *testing.T) {
	p := ProblemOne{}

	err := p.SolveSecond()

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblemOne(t *testing.T) {
	p := ProblemOne{}

	err := p.Solve()

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}
