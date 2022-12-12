package day11

import (
	"fmt"
	"testing"
)

var testFile string = "problem-11-example.txt"
var realFile string = "problem-11.txt"

// func TestSolveFirst(t *testing.T) {

// 	expected := 10605
// 	actual := SolveFirst(testFile)

// 	if expected != actual {
// 		t.Errorf("Expected %d but got %d", expected, actual)
// 		t.Fail()
// 	}
// 	answer := SolveFirst(realFile)
// 	fmt.Printf("\n-----\n --ANSWER Day 11 part 1 = %d \n-----\n", answer)
// }

func TestSolveSecond(t *testing.T) {

	expected := 24
	actual := SolveSecond(testFile, 1)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}

	expected = 10197
	actual = SolveSecond(testFile, 20)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}

	expected = 27019168
	actual = SolveSecond(testFile, 1000)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}

	expected = 2713310158
	actual = SolveSecond(testFile, 10000)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}

	expected = 0
	actual = SolveSecond(realFile, 10000)

	if expected == actual {
		t.Errorf("Did not expect %d", expected)
		t.Fail()
	}
	fmt.Printf("\n-----\n --ANSWER Day 11 part 2 = %d \n-----\n", actual)

}

func BenchmarkSolveFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveFirst(realFile)
	}
}

func BenchmarkSolveSecond(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolveSecond(realFile, 10000)
	}
}
