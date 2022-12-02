package day02

import (
	"fmt"
	"testing"
)

var testFile string = "problem-2-example.txt"
var realFile string = "problem-2.txt"

func TestSolveFirst(t *testing.T) {
	expected := 15
	actual := SolveFirst(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveFirst(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 02 part 1 = %d \n-----\n", answer)
}

func TestSolveSecondA(t *testing.T) {
	expected := 12
	actual := SolveSecondA(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveSecondA(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 02 Part 2A = %d \n-----\n", answer)
}
func TestSolveSecondB(t *testing.T) {
	expected := 12
	actual := SolveSecondB(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveSecondB(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 02 part 2B = %d \n-----\n", answer)
}

func BenchmarkSolveFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveFirst(realFile)
	}
}
func BenchmarkSolveSecondA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveSecondA(realFile)
	}
}
func BenchmarkSolveSecondB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveSecondB(realFile)
	}
}
