package problems

import (
	"testing"
)

var problem2TestFile string = "problem-2-example.txt"
var problem2RealFile string = "problem-2.txt"

func TestProblem2PartOne(t *testing.T) {

	var p Problem = ProblemTwo{}
	err := p.SolveFirst(problem2TestFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblem2PartOneReal(t *testing.T) {
	var p Problem = ProblemTwo{}
	err := p.SolveFirst(problem2RealFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblem2PartTwo(t *testing.T) {
	var p Problem = ProblemTwo{}
	err := p.SolveFirst(problem2TestFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblem2PartTwoReal(t *testing.T) {
	var p Problem = ProblemTwo{}
	err := p.SolveFirst(problem2RealFile)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblem2(t *testing.T) {
	var p Problem = ProblemTwo{}
	err := p.SolveFirst(problem2TestFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}

func TestProblem2Real(t *testing.T) {
	var p Problem = ProblemTwo{}
	err := p.SolveFirst(problem2RealFile)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}
