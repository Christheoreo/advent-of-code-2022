package day11

import (
	"fmt"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

type Monkey struct {
	Name            string
	WorryLevels     []int
	Operation       byte
	TestOperation   int
	MonkeyToThrowTo [2]int
}

func SolveFirst(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	monkeys := sortData(data)
	total := 0
	for _, monkey := range monkeys {
		// fmt.Printf("Monkey %s worry lvels = %v\n", monkey.Name, monkey.WorryLevels)
		fmt.Printf("Monkey %+v\n\n", monkey)
	}
	return total
}

func SolveSecond(filename string) int {
	// data, _ := filereader.ReadFileToStringArray(filename)
	x := 1 // middle spriate pos

	return x
}
