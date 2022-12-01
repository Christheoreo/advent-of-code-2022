package problems

import "testing"

var problem1TestFile string = "problem-1-example.txt"
var problem1RealFile string = "problem-1.txt"

func TestProblem1PartOne(t *testing.T) {
	p := ProblemOne{}
	err := p.SolveFirst(problem1TestFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblem1PartOneReal(t *testing.T) {
	p := ProblemOne{}
	err := p.SolveFirst(problem1RealFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblem1PartTwo(t *testing.T) {
	p := ProblemOne{}
	err := p.SolveFirst(problem1TestFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblem1PartTwoReal(t *testing.T) {
	p := ProblemOne{}
	err := p.SolveFirst(problem1RealFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblem1(t *testing.T) {
	p := ProblemOne{}
	err := p.SolveFirst(problem1TestFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblem1Real(t *testing.T) {
	p := ProblemOne{}
	err := p.SolveFirst(problem1RealFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}
