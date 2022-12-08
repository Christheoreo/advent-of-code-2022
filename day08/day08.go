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
	rowLength := len(data[0])

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
			above := make([]int, rowIndex+1)
			below := make([]int, rowsTotal-1-rowIndex+1)
			for k, l := rowIndex, 0; k >= 0; k, l = k-1, l+1 {
				above[l] = rows[k][i]
			}
			for k, l := rowIndex, 0; k < rowsTotal; k, l = k+1, l+1 {
				below[l] = rows[k][i]
			}
			a, b, c, d := findLeft(row, i, tree), findRight(row, i, rowLength, tree), findNext(above, 0, tree), findNext(below, 0, tree)
			score := a * b * c * d
			if score > highscore {
				highscore = score
			}
		}
	}

	return highscore
}

func findRight(row []int, c int, rowLength int, tree int) int {
	count := 1
	// is there any to the right?
	if c > rowLength-2 {
		return 0
	}
	c += 1
	newTree := row[c]
	if tree > newTree {
		count += findRight(row, c, rowLength, tree)
	}
	return count
}

func findLeft(row []int, c int, tree int) int {
	count := 1
	// is there any to the right?
	if c < 1 {
		return 0
	}
	c -= 1
	newTree := row[c]
	if tree > newTree {
		count += findLeft(row, c, tree)
	}
	return count
}

func findNext(col []int, r int, tree int) int {
	count := 1
	if r > len(col)-2 {
		return 0
	}
	r += 1
	newTree := col[r]
	if tree > newTree {
		count += findNext(col, r, tree)
	}
	return count
}
