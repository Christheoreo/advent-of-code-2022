package day05

import (
	"fmt"
	"testing"
)

var testFile string = "problem-5-example.txt"
var realFile string = "problem-5.txt"

func TestSolveFirst(t *testing.T) {
	expected := "CMZ"
	actual := SolveFirst(testFile)
	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
	// answer := SolveFirst(realFile)
	// fmt.Printf("\n-----\n --ANSWER Day 05 part 1 = %s \n-----\n", answer)
}

func TestSolveSecond(t *testing.T) {
	expected := ""
	actual := SolveSecond(testFile)
	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
	answer := SolveSecond(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 05 part 2 = %s \n-----\n", answer)
}

func BenchmarkSolveFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveFirst(realFile)
	}
}
func BenchmarkSolveSecond(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveSecond(realFile)
	}
}
