package problems

import "testing"

var problemXTestFile string = "problem-X-example.txt"
var problemXRealFile string = "problem-X.txt"

func TestProblemXPartOne(t *testing.T) {
	p := ProblemOne{}
	err := p.SolveFirst(problemXTestFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblemXPartOneReal(t *testing.T) {
	p := ProblemOne{}
	err := p.SolveFirst(problemXRealFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblemXPartTwo(t *testing.T) {
	p := ProblemOne{}
	err := p.SolveFirst(problemXTestFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblemXPartTwoReal(t *testing.T) {
	p := ProblemOne{}
	err := p.SolveFirst(problemXRealFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblemX(t *testing.T) {
	p := ProblemOne{}
	err := p.SolveFirst(problemXTestFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblemXReal(t *testing.T) {
	p := ProblemOne{}
	err := p.SolveFirst(problemXRealFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}
