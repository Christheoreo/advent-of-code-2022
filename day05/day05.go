package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

func SolveFirst(filename string) string {
	data, _ := filereader.ReadFileToStringArray(filename)
	answer := ""

	a := make([][]string, 0)

	m := map[int][]string{}
	// get the column names
	for _, val := range data {
		if strings.Contains(val, "move") || strings.Contains(val, "[") {
			continue
		}

		parts := strings.Split(val, " ")

		for _, part := range parts {
			if part == "" {
				continue
			}
			num, _ := strconv.Atoi(part)
			m[num] = make([]string, 0)
		}
	}
	for i, val := range data {
		arr := strings.Split(val, " ")
		a = append(a, make([]string, 0))

		if !strings.Contains(val, "move") {
			if !strings.Contains(val, "[") {
				continue
			}

			for _, character := range arr {
				if character == "" || character == "[" || character == "]" {
					continue
				}

				a[i] = append(a[i], character)
			}
			continue
		}

		// fmt.Println(arr)

		// instructions time
		instructions := [3]int{0, 0, 0}
		index := 0
		for _, character := range arr {
			if character == "move" || character == "from" || character == "to" {
				continue
			}

			num, _ := strconv.Atoi(character)

			instructions[index] = num
			index++
		}

		fmt.Println(instructions)

	}
	fmt.Println(a)

	return answer
}

func SolveSecond(filename string) string {
	// data, _ := filereader.ReadFileToStringArray(filename)
	pairs := ""

	return pairs
}
