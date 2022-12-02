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
func TestSolveSecondAttempt2(t *testing.T) {
	expected := 12
	actual := SolveSecondAttempt2(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveSecondAttempt2(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 02 part 2B = %d \n-----\n", answer)
}
func TestSolveSecondAttempt3(t *testing.T) {
	expected := 12
	actual := SolveSecondAttempt3(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveSecondAttempt3(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 02 part 2C = %d \n-----\n", answer)
}

func TestSolveSecondAttempt4(t *testing.T) {
	expected := 12
	actual := SolveSecondAttempt4(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveSecondAttempt4(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 02 part 2D = %d \n-----\n", answer)
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

func BenchmarkSolveSecondAttempt3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveSecondAttempt3(realFile)
	}
}

func BenchmarkSolveSecondAttempt4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveSecondAttempt4(realFile)
	}
}
