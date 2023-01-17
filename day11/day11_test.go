package day11

import (
	"testing"
)

var testFile string = "problem-11-example.txt"
var realFile string = "problem-11.txt"

func TestSolveFirst(t *testing.T) {
	expected := 13140
	actual := SolveFirst(testFile)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}
	// answer := SolveFirst(realFile)
	// fmt.Printf("\n-----\n --ANSWER Day 11 part 1 = %d \n-----\n", answer)
}

// func TestSolveSecond(t *testing.T) {
// 	expected := 0
// 	actual := SolveSecond(testFile)
// 	if expected != actual {
// 		t.Errorf("Expected \n%d\n but got \n%d\n", expected, actual)
// 		t.Fail()
// 	}
// 	answer := SolveSecond(realFile)

// 	fmt.Printf("\n-----\n --ANSWER Day 11 part 2 = \n%d\n \n-----\n", answer)
// }

func BenchmarkSolveFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveFirst(realFile)
	}
}

// func BenchmarkSolveSecond(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		SolveSecond(realFile)
// 	}
// }
