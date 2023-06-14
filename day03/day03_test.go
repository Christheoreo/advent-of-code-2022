package day03

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed real.txt
var data string

//go:embed example.txt
var testData string

func TestSolveFirst(t *testing.T) {
	expected := 157
	actual := SolveFirst(testData)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}

	answer := SolveFirst(data)
	fmt.Printf("ANSWER Day 03 part 1 = %d\n", answer)
}

func TestSolveSecond(t *testing.T) {
	expected := 70
	actual := SolveSecond(testData)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}

	answer := SolveSecond(data)
	fmt.Printf("ANSWER Day 03 part 2 = %d\n", answer)
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
