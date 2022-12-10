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

func TestSolveSecond(t *testing.T) {
	expected := [6]string{
		"##..##..##..##..##..##..##..##..##..##..",
		"###...###...###...###...###...###...###.",
		"####....####....####....####....####....",
		"#####.....#####.....#####.....#####.....",
		"######......######......######......####",
		"#######.......#######.......#######.....",
	}
	actual := SolveSecond(testFile)
	if expected != actual {
		t.Errorf("Expected \n%v\n but got \n%v\n", expected, actual)
		t.Fail()
	}
	answer := SolveSecond(realFile)

	formattedAnswer := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", answer[0], answer[1], answer[2], answer[3], answer[4], answer[5])
	fmt.Printf("\n-----\n --ANSWER Day 10 part 2 = \n%s\n \n-----\n", formattedAnswer)
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
