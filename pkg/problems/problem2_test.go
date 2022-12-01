package problems

import "testing"

func TestProblem2PartOne(t *testing.T) {
	p := ProblemTwo{}

	err := p.SolveFirst()

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblem2PartTwo(t *testing.T) {
	p := ProblemTwo{}

	err := p.SolveSecond()

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblem2(t *testing.T) {
	p := ProblemTwo{}

	err := p.Solve()

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}
