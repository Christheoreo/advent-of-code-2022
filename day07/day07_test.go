package day07

import (
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
	// answer := SolveFirst(realFile)
	// fmt.Printf("\n-----\n --ANSWER Day 07 part 1 = %s \n-----\n", answer)
}

// func TestSolveSecond(t *testing.T) {
// 	expected := 0
// 	actual := SolveSecond(testFile)
// 	if expected != actual {
// 		t.Errorf("Expected %s but got %s", expected, actual)
// 		t.Fail()
// 	}
// 	answer := SolveSecond(realFile)
// 	fmt.Printf("\n-----\n --ANSWER Day 07 part 2 = %s \n-----\n", answer)
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
