package day08

import (
	"strconv"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

func SolveFirst(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	count := 0
	rowsTotal := len(data)
	rows := make([][]int, rowsTotal)
	for rowIndex, line := range data {
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			rows[rowIndex] = append(rows[rowIndex], num)
		}
	}
	for rowIndex := 1; rowIndex < len(rows[0])-1; rowIndex++ {
		row := rows[rowIndex]

		for col := 1; col < len(row)-1; col++ {
			tree := row[col]
			leftSide := row[:col]
			rightSide := row[col+1:]
			above := make([]int, 0)
			below := make([]int, 0)

			for r := rowIndex - 1; r >= 0; r-- {
				above = append(above, rows[r][col])
			}
			for r := rowIndex + 1; r <= len(rows)-1; r++ {
				below = append(below, rows[r][col])
			}

			slices := [4][]int{rightSide, leftSide, above, below}
			visible := true

			for _, slice := range slices {
				visible = true

				for _, num := range slice {
					if num >= tree {
						visible = false
						break
					}
				}
				if !visible {
					continue
				}
				count++
				break
			}
		}
	}
	count += (len(rows) * 2) + ((len(rows[0]) - 2) * 2)
	return count
}
func SolveSecond(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	highscore := 0
	rowsTotal := len(data)
	rows := make([][]int, rowsTotal)

	// setup the ints
	for rowIndex, line := range data {
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			rows[rowIndex] = append(rows[rowIndex], num)
		}
	}

	for rowIndex := 1; rowIndex < rowsTotal-1; rowIndex++ {
		row := rows[rowIndex]
		for i := 1; i < len(row)-1; i++ {
			tree := row[i]

			rightSide := row[i+1:]
			leftSide := make([]int, 0)
			above := make([]int, 0)
			below := make([]int, 0)

			for rI := rowIndex - 1; rI >= 0; rI-- {
				above = append(above, rows[rI][i])
			}
			for rI := rowIndex + 1; rI <= rowsTotal-1; rI++ {
				below = append(below, rows[rI][i])
			}
			for cI := i - 1; cI >= 0; cI-- {
				leftSide = append(leftSide, row[cI])
			}

			numbers := [4]int{0, 0, 0, 0}
			slices := [4][]int{rightSide, leftSide, above, below}
			for sliceI, slice := range slices {
				n := 0
				for _, num := range slice {
					n++
					if tree > num {
						continue
					}
					break
				}
				numbers[sliceI] = n
			}

			score := numbers[0] * numbers[1] * numbers[2] * numbers[3]
			if score > highscore {
				highscore = score
			}
		}
	}

	return highscore
}
