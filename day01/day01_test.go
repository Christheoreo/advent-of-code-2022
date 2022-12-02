package day01

import (
	"fmt"
	"testing"
)

var testFile string = "problem-1-example.txt"
var realFile string = "problem-1.txt"

func TestSolveFirst(t *testing.T) {
	expected := 24000
	actual := SolveFirst(testFile)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}

	answer := SolveFirst(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 01 part 1 = %d \n-----\n", answer)
}

func TestSolveSecond(t *testing.T) {
	expected := 45000
	actual := SolveSecond(testFile)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}

	answer := SolveFirst(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 01 part 2 = %d \n-----\n", answer)
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
