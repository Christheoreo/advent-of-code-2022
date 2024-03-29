package day05

import (
	"strconv"
	"strings"
)

const emptyS = byte(32)

func sortColumns(data []string) [][]byte {
	mapping := make([][]byte, 0)
	for rowIndex := len(data) - 2; rowIndex >= 0; rowIndex-- {
		val := data[rowIndex]
		valLength := len(val)

		for i, j := 1, 1; i <= valLength; i, j = i+4, j+1 {

			r := val[i]

			if r == emptyS {
				continue
			}

			ok := len(mapping) >= j

			if !ok {
				for i := len(mapping); i <= j-1; i++ {
					mapping = append(mapping, make([]byte, 0))
				}
			}
			mapping[j-1] = append(mapping[j-1], r)
		}
	}
	return mapping
}

func getData(data string) (string, string) {
	tippingPoint := strings.Index(data, "\nmove")
	columnsData := data[:tippingPoint-1]
	instructionsData := data[tippingPoint+1:]
	return columnsData, instructionsData
}

func getAnswerFromMapping(mapping [][]byte) string {
	answer := ""
	for _, val := range mapping {
		answer += string(val[len(val)-1])
	}
	return answer
}

func SolveFirst(data string) string {
	columnsData, instructionsData := getData(data)
	numbersSplit := strings.Split(columnsData, "\n")
	mapping := sortColumns(numbersSplit)

	for _, val := range strings.Split(instructionsData, "\n") {
		s := strings.Split(val, " ")
		move, _ := strconv.Atoi(s[1])
		from, _ := strconv.Atoi(s[3])
		to, _ := strconv.Atoi(s[5])

		for i := 1; i <= move; i++ {
			f := from - 1
			t := to - 1
			lastEle := mapping[f][len(mapping[f])-1]
			mapping[f] = mapping[f][:len(mapping[f])-1]
			mapping[t] = append(mapping[t], lastEle)
		}
	}

	return getAnswerFromMapping(mapping)
}

func SolveSecond(data string) string {
	columnsData, instructionsData := getData(data)
	numbersSplit := strings.Split(columnsData, "\n")
	mapping := sortColumns(numbersSplit)
	for _, val := range strings.Split(instructionsData, "\n") {
		s := strings.Split(val, " ")

		move, _ := strconv.Atoi(s[1])
		from, _ := strconv.Atoi(s[3])
		to, _ := strconv.Atoi(s[5])

		m := move
		f := from - 1
		t := to - 1
		formula := len(mapping[f]) - m
		toMove := mapping[f][formula:]
		mapping[f] = mapping[f][:formula]
		mapping[t] = append(mapping[t], toMove...)

	}

	return getAnswerFromMapping(mapping)
}
