package day09

import (
	"fmt"
	"strconv"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

func SolveSecond(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	instructions := make([]Instruction, len(data))
	positions := make(map[string]int)
	positions["x0y0"] = 1

	snake := [10][2]int{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	for index, val := range data {
		num, _ := strconv.Atoi(val[2:])
		instruction := Instruction{
			Direction: string(val[0]),
			Count:     num,
		}
		instructions[index] = instruction
	}

	directionMap := map[string]int{
		"U": 1,
		"D": -1,
		"L": -1,
		"R": 1,
	}

	for _, instruction := range instructions {

		for i := 0; i < instruction.Count; i++ {

			if instruction.Direction == "L" || instruction.Direction == "R" {
				snake[0][0] += directionMap[instruction.Direction]
			} else {
				snake[0][1] += directionMap[instruction.Direction]
			}

			for snakeIndex := 1; snakeIndex < len(snake); snakeIndex++ {
				h := snake[snakeIndex-1]
				t := snake[snakeIndex]

				reachable := isInReach(h, t)
				if reachable {
					break
				}
				// is it in the same Y Axis?
				if t[1] == h[1] {
					if h[0] > t[0] {
						snake[snakeIndex][0] += 1
					} else {
						snake[snakeIndex][0] -= 1
					}
					continue
				}

				// Is it in the same X Axis?
				if t[0] == h[0] {
					if h[1] > t[1] {
						snake[snakeIndex][1] += 1
					} else {
						snake[snakeIndex][1] -= 1
					}
					continue
				}

				// if we get here, we need to make a diagonal step

				if h[0] > t[0] {
					snake[snakeIndex][0] += 1
				} else {
					snake[snakeIndex][0] -= 1
				}

				if h[1] > t[1] {
					snake[snakeIndex][1] += 1
				} else {
					snake[snakeIndex][1] -= 1
				}

			}

			key := fmt.Sprintf("x%dy%d", snake[9][0], snake[9][1])
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
