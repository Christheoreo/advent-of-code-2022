package day08

import (
	"fmt"
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
			fmt.Printf("row =%d col = %d riughtise = %v\n", rowIndex, col, rightSide)

			visible := true
			for _, num := range leftSide {
				if num >= tree {
					visible = false
					break
				}
			}
			if visible {
				count++
				fmt.Printf("Tree  %d (%d, %d) is visible!\n", tree, rowIndex, col)
				continue
			}
			visible = true

			for _, num := range rightSide {
				if num >= tree {
					visible = false
					break
				}
			}
			if visible {
				count++
				fmt.Printf("Tree  %d (%d, %d) is visible!\n", tree, rowIndex, col)
				continue
			}
			visible = true

			above := make([]int, 0)
			below := make([]int, 0)

			if rowIndex > 0 {
				for r := rowIndex - 1; r >= 0; r-- {
					above = append(above, rows[r][col])
				}
			}

			if rowIndex < len(rows)-1 {
				for r := rowIndex + 1; r <= len(rows)-1; r++ {
					below = append(below, rows[r][col])
				}
			}

			for _, num := range above {
				if num >= tree {
					visible = false
					break
				}
			}
			if visible {
				count++
				fmt.Printf("Tree  %d (%d, %d) is visible!\n", tree, rowIndex, col)
				continue
			}
			visible = true
			for _, num := range below {
				if num >= tree {
					visible = false
					break
				}
			}
			if visible {
				count++
				fmt.Printf("Tree  %d (%d, %d) is visible!\n", tree, rowIndex, col)
				continue
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

	for rowIndex, line := range data {
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			rows[rowIndex] = append(rows[rowIndex], num)
		}
	}

	fmt.Println(rows)
	for rowIndex := 1; rowIndex < rowsTotal-1; rowIndex++ {

		row := rows[rowIndex]
		fmt.Println(row)

		for i := 1; i < len(row)-1; i++ {
			tree := row[i]

			leftSide := row[:i]
			// reverse left side
			// for i, j := 0, len(leftSide)-1; i < j; i, j = i+1, j-1 {
			// 	leftSide[i], leftSide[j] = leftSide[j], leftSide[i]
			// }
			rightSide := row[i+1:]

			above := make([]int, 0)
			below := make([]int, 0)

			for rI := rowIndex - 1; rI >= 0; rI-- {
				above = append(above, rows[rI][i])
			}
			for rI := rowIndex + 1; rI <= rowsTotal-1; rI++ {
				below = append(below, rows[rI][i])
			}

			slices := [3][]int{rightSide, above, below}

			numbers := [4]int{0, 0, 0, 0}
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

			// left  side needs to be done backwards
			n := 0
			for i := len(leftSide) - 1; i >= 0; i-- {
				n++
				if tree > leftSide[i] {
					continue
				}
				break
			}
			numbers[3] = n

			score := numbers[0] * numbers[1] * numbers[2] * numbers[3]

			fmt.Printf("for tree %d in {%d,%d} we had %v which score %d\n", tree, rowIndex, i, numbers, score)

			if score > highscore {
				highscore = score
			}

			fmt.Println(tree)
			fmt.Println(leftSide)
			fmt.Println(rightSide)
			fmt.Println(above)
			fmt.Println(below)
			fmt.Println()
			fmt.Println()

		}
	}

	return highscore
}
