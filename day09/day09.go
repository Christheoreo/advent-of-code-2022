package day09

import (
	"fmt"
	"strconv"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

type Direction int64

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Instruction struct {
	// Direction Direction
	Direction string
	Count     int
}

func SolveFirst(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	instructions := make([]Instruction, len(data))

	for index, val := range data {
		num, _ := strconv.Atoi(val[2:])
		instruction := Instruction{
			Direction: string(val[0]),
			Count:     num,
		}
		instructions[index] = instruction
	}

	for _, instuction := range instructions {
		fmt.Printf("Direction = %s and count = %d\n", instuction.Direction, instuction.Count)
	}
	return 0
}
func SolveSecond(filename string) int {
	// data, _ := filereader.ReadFileToStringArray(filename)
	return 0
}
