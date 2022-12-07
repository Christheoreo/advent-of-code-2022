package day07

import (
	"fmt"
	"testing"
)

var testFile string = "problem-7-example.txt"
var realFile string = "problem-7.txt"

func TestSolveFirst(t *testing.T) {
	expected := 95437
	actual := SolveFirst(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveFirst(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 07 part 1 = %d \n-----\n", answer)
}

func TestSolveSecond(t *testing.T) {
	expected := 24933642
	actual := SolveSecond(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveSecond(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 07 part 2 = %d \n-----\n", answer)
}

func TestSolveFirstA(t *testing.T) {
	expected := 95437
	actual := SolveFirstA(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveFirstA(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 07 part 1 = %d \n-----\n", answer)
}

func TestSolveSecondA(t *testing.T) {
	expected := 24933642
	actual := SolveSecondA(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveSecondA(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 07 part 2 = %d \n-----\n", answer)
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

func BenchmarkSolveFirstA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveFirstA(realFile)
	}
}
func BenchmarkSolveSecondA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveSecondA(realFile)
	}
}
