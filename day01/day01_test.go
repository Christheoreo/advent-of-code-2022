package day01

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed problem-1.txt
var data string

//go:embed problem-1-example.txt
var testData string

func TestSolveFirst(t *testing.T) {
	expected := 24000
	actual := SolveFirst(testData)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}

	answer := SolveFirst(data)
	fmt.Printf("\n-----\n --ANSWER Day 01 part 1 = %d \n-----\n", answer)
}

func TestSolveSecond(t *testing.T) {
	expected := 45000
	actual := SolveSecond(testData)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}

	answer := SolveSecond(data)
	fmt.Printf("\n-----\n --ANSWER Day 01 part 2 = %d \n-----\n", answer)
}

func BenchmarkSolveFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveFirst(data)
	}
}

func BenchmarkSolveSecond(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveSecond(data)
	}
}
