package day08

import (
	"fmt"
	"testing"
)

var testFile string = "problem-8-example.txt"
var realFile string = "problem-8.txt"

func TestSolveFirst(t *testing.T) {
	expected := 21
	actual := SolveFirst(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveFirst(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 08 part 1 = %d \n-----\n", answer)
}

func TestSolveFirstFaster(t *testing.T) {
	expected := 21
	actual := SolveFirstFaster(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveFirstFaster(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 08 part 1 = %d \n-----\n", answer)
}

func TestSolveFirstB(t *testing.T) {
	expected := 21
	actual := SolveFirstB(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveFirstB(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 08 part 1 = %d \n-----\n", answer)
}

func TestSolveSecond(t *testing.T) {
	expected := 8
	actual := SolveSecond(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveSecond(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 08 part 2 = %d \n-----\n", answer)
}

func TestSolveSecondB(t *testing.T) {
	expected := 8
	actual := SolveSecondFaster(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveSecond(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 08 part 2 = %d \n-----\n", answer)
}

func BenchmarkSolveFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveFirst(realFile)
	}
}

func BenchmarkSolveFirstFaster(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveFirstFaster(realFile)
	}
}

func BenchmarkSolveFirstB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveFirstB(realFile)
	}
}

func BenchmarkSolveSecond(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveSecond(realFile)
	}
}

func BenchmarkSolveSecondFaster(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveSecondFaster(realFile)
	}
}
