package day05

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
	expected := "CMZ"
	actual := SolveFirst(testData)

	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
		return
	}

	answer := SolveFirst(data)
	fmt.Printf("ANSWER Day 05 part 1 = %s\n", answer)
}

// func TestSolveSecond(t *testing.T) {
// 	expected := 4
// 	actual := SolveSecond(testData)

// 	if expected != actual {
// 		t.Errorf("Expected %d but got %d", expected, actual)
// 		t.Fail()
// 	}

// 	answer := SolveSecond(data)
// 	fmt.Printf("ANSWER Day 04 part 2 = %d\n", answer)
// }

// func BenchmarkSolveFirst(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		SolveFirst(data)
// 	}
// }

// func BenchmarkSolveSecond(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		SolveSecond(data)
// 	}
// }
