package day10

import (
	"fmt"
	"testing"
)

var testFile string = "problem-10-example.txt"
var realFile string = "problem-10.txt"

func TestSolveFirst(t *testing.T) {
	expected := 13140
	actual := SolveFirst(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	answer := SolveFirst(realFile)
	fmt.Printf("\n-----\n --ANSWER Day 10 part 1 = %d \n-----\n", answer)
}

// func TestSolveSecond(t *testing.T) {
// 	expected := 0
// 	actual := SolveSecond(testFile)
// 	if expected != actual {
// 		t.Errorf("Expected %d but got %d", expected, actual)
// 		t.Fail()
// 	}
// 	answer := SolveSecond(realFile)
// 	fmt.Printf("\n-----\n --ANSWER Day 10 part 2 = %d \n-----\n", answer)
// }

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