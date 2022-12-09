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
	positions := make(map[string]int)
	positions["x0y0"] = 1
	h := [2]int{0, 0}
	t := [2]int{0, 0}

	for index, val := range data {
		num, _ := strconv.Atoi(val[2:])
		instruction := Instruction{
			Direction: string(val[0]),
			Count:     num,
		}
		instructions[index] = instruction
	}

	for _, instruction := range instructions {

		for i := 0; i < instruction.Count; i++ {
			if instruction.Direction == "L" {
				// Go Left
				h[0] -= 1
				if isInReach(h, t) {
					continue
				}
				t[0] = h[0] + 1
				t[1] = h[1]

			} else if instruction.Direction == "R" {
				// Go Right
				h[0] += 1
				if isInReach(h, t) {
					continue
				}
				t[0] = h[0] - 1
				t[1] = h[1]

			} else if instruction.Direction == "U" {
				// Go Up
				h[1] += 1
				if isInReach(h, t) {
					continue
				}
				t[1] = h[1] - 1
				t[0] = h[0]

			} else {
				// Go Down
				h[1] -= 1

				if isInReach(h, t) {
					continue
				}
				t[1] = h[1] + 1
				t[0] = h[0]
			}

			key := fmt.Sprintf("x%dy%d", t[0], t[1])
			_, ok := positions[key]
			if !ok {
				positions[key] = 1
				continue
			}
			positions[key]++
		}
	}
	return len(positions)
}
func SolveSecond(filename string) int {
	// data, _ := filereader.ReadFileToStringArray(filename)
	return 0
}

func isInReach(head [2]int, tail [2]int) bool {
	xDiff := head[0] - tail[0]
	yDiff := head[1] - tail[1]

	return xDiff > -2 && xDiff < 2 && yDiff > -2 && yDiff < 2
}
