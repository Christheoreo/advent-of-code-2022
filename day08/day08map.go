package day08

import (
	"strconv"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

func SolveFirstB(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	count := 0
	m := make(map[int]map[int]int)
	rowsTotal := len(data)
	colsTotal := len(data[0])

	for rowIndex, line := range data {
		for colIndex, char := range line {
			num, _ := strconv.Atoi(string(char))
			_, ok := m[rowIndex]
			if !ok {
				m[rowIndex] = make(map[int]int)
			}
			m[rowIndex][colIndex] = num
		}
	}

	for r := 1; r < rowsTotal-1; r++ {
		for c := 1; c < colsTotal-1; c++ {
			// Can it be seen from the left?
			tree := m[r][c]
			canBeSeen := true
			for i := c - 1; i >= 0; i-- {
				val := m[r][i]
				if val >= tree {
					canBeSeen = false
					break
				}
			}
			if canBeSeen {
				count++
				// fmt.Printf("visible tree forund %d {%d, %d} from left\n", tree, r, c)
				continue
			}
			canBeSeen = true
			// can it be seen from the right?
			for i := c + 1; i < colsTotal; i++ {
				val := m[r][i]
				if val >= tree {
					canBeSeen = false
					break
				}
			}
			if canBeSeen {
				count++
				// fmt.Printf("visible tree forund %d {%d, %d} from right\n", tree, r, c)
				continue
			}
			canBeSeen = true

			// can it be seen from above?
			for i := r - 1; i >= 0; i-- {
				val := m[i][c]
				if val >= tree {
					canBeSeen = false
					break
				}
			}
			if canBeSeen {
				count++
				// fmt.Printf("visible tree forund %d {%d, %d} from above\n", tree, r, c)
				continue
			}
			canBeSeen = true
			// can it be seen from below?
			for i := r + 1; i < rowsTotal; i++ {
				val := m[i][c]
				if val >= tree {
					canBeSeen = false
					break
				}
			}
			if !canBeSeen {
				continue
			}
			count++
		}
	}

	return count + (len(data) * 2) + ((len(data[0]) - 2) * 2)
}
