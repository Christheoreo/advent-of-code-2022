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

func TestSolveSecond(t *testing.T) {
	expected := 12
	actual := SolveSecond(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveSecond(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 02 = %d \n-----\n", answer)
}
func TestSolveSecondAttempt(t *testing.T) {
	expected := 12
	actual := SolveSecondAttempt2(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveSecondAttempt2(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 02 part 2B = %d \n-----\n", answer)
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
func BenchmarkSolveSecondAttempt2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveSecondAttempt2(realFile)
	}
}
