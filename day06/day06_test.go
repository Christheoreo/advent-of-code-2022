package day06

import (
	"fmt"
	"testing"
)

var testFile string = "problem-6-example.txt"
var testFileA string = "problem-6-examplea.txt"
var testFileB string = "problem-6-exampleb.txt"
var testFileC string = "problem-6-examplec.txt"
var testFileD string = "problem-6-exampled.txt"
var realFile string = "problem-6.txt"

func TestSolveFirst(t *testing.T) {
	m := map[string]int{
		testFile:  7,
		testFileA: 5,
		testFileB: 6,
		testFileC: 10,
		testFileD: 11,
	}
	for key, val := range m {
		expected := val
		actual := SolveFirst(key)

		if expected != actual {
			t.Errorf("Expected %d but got %d", expected, actual)
			t.Fail()
		}
	}

	answer := SolveFirst(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 06 part 1 = %d \n-----\n", answer)
}

func TestSolveSecond(t *testing.T) {
	m := map[string]int{
		testFile:  19,
		testFileA: 23,
		testFileB: 23,
		testFileC: 29,
		testFileD: 26,
	}
	for key, val := range m {
		expected := val
		actual := SolveSecond(key)

		if expected != actual {
			t.Errorf("Expected %d but got %d", expected, actual)
			t.Fail()
		}
	}

	answer := SolveSecond(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 06 part 2 = %d \n-----\n", answer)
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
