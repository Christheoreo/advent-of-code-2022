package day01

import (
	"fmt"
	"testing"
	"time"

	"github.com/christheoreo/advent-of-code-2022/internal/timetrack"
)

var testFile string = "problem-1-example.txt"
var realFile string = "problem-1.txt"

func TestSolveFirst(t *testing.T) {
	expected := 24000
	test := time.Now()

	actual := SolveFirst(testFile)
	timetrack.TimeTrack(test, "Day 01 part 01 (test data)")

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}

	real := time.Now()
	answer := SolveFirst(realFile)
	timetrack.TimeTrack(real, "Day 01 part 01")
	fmt.Printf("Day 01 part 1 = %d\n", answer)
}

func TestSolveSecond(t *testing.T) {
	expected := 45000
	test := time.Now()

	actual := SolveSecond(testFile)
	timetrack.TimeTrack(test, "Day 01 part 02 (test data)")

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}

	real := time.Now()
	answer := SolveFirst(realFile)
	timetrack.TimeTrack(real, "Day 01 part 02")
	fmt.Printf("Day 01 part 2 = %d\n", answer)
}
