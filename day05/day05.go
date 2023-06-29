package day05

import (
	"fmt"
	"strconv"
	"strings"
)

// func SolveFirst(data string) string {
// 	// Find the first line including " 1"
// 	lines := strings.Split(data, "\n")
// 	numberIndex := 0
// 	for lineIndex, line := range lines {
// 		if strings.Contains(line, " 1") {
// 			numberIndex = lineIndex
// 			break
// 		}
// 	}

// 	numberLine := lines[numberIndex]

// 	x := strings.TrimSpace(numberLine)
// 	columnCount := len(strings.Split(x, " "))

// 	return ""
// }

const emptyS = byte(32)

func SolveFirst(data string) string {
	// Find the first line including " 1"
	// mapping := map[int][]byte{}
	mapping := make([][]byte, 0)
	tippingPoint := strings.Index(data, "\nmove")
	// lastRow := tippingPoint - 2

	numbers := data[:tippingPoint-1]
	// fmt.Println((numbers))
	other := data[tippingPoint+1:]
	// columnCount, _ := strconv.Atoi(string(data[tippingPoint-3]))
	// fmt.Println(columnCount)

	// 4 * col - 1 = index

	// Start at index 1, then add 4 to get col value until we run out of string
	numbersSplit := strings.Split(numbers, "\n")

	// fmt.Println(numbers)
	// fmt.Println(numbersSplit)
	// for _, val := range numbersSplit {
	for xx := len(numbersSplit) - 2; xx >= 0; xx-- {
		val := numbersSplit[xx]
		l := len(val)
		// fmt.Printf("Val = %s and l = %d\n", val, l)

		for i, j := 1, 1; i <= l; i, j = i+4, j+1 {

			r := val[i]

			if r == emptyS {
				continue
			}

			ok := len(mapping) >= j
			// _, ok := mapping[j]

			if !ok {
				// fmt.Printf("here - %d is greater than %d\n", j, len(mapping))
				for i := len(mapping); i <= j-1; i++ {
					mapping = append(mapping, make([]byte, 0))
				}
				// mapping[j] = make([]byte, 0)
			}
			// fmt.Println(len(mapping))
			// fmt.Println(j)
			mapping[j-1] = append(mapping[j-1], r)
		}
	}

	for col, v := range mapping {
		fmt.Printf("Length of %d is %d\n", col, len(v))
		for index, x := range v {
			fmt.Printf("Col %d index = %d = %s\n", col, index, string(x))
		}
	}

	for _, val := range strings.Split(other, "\n") {
		s := strings.Split(val, " ")

		// move, from, to := 0,0,0
		move, _ := strconv.Atoi(s[1])
		from, _ := strconv.Atoi(s[3])
		to, _ := strconv.Atoi(s[5])

		fmt.Println("abc")

		for i := 1; i <= move; i++ {
			f := from - 1
			t := to - 1
			lastEle := mapping[f][len(mapping[f])-1]
			mapping[f] = mapping[f][:len(mapping[f])-1]
			mapping[t] = append(mapping[t], lastEle)

			// lastEle := mapping[from][0]
			// if len(mapping[from]) == 1 {
			// 	mapping[from] = make([]byte, 0)
			// } else {
			// 	mapping[from] = mapping[from][1:]
			// }
			// mapping[to] = append([]byte{lastEle}, mapping[to]...)
		}

		// fmt.Printf("%s - %s - %s\n", s[1], s[3], s[5])
	}

	answer := ""
	for _, val := range mapping {
		answer += string(val[len(val)-1])
	}

	// for k := 1; k <= len(mapping); k++ {
	// 	val := mapping[k]
	// 	answer += string(val[len(val)-1])
	// }

	// fmt.Println(other)

	return answer
}

func SolveSecond(data string) string {
	return ""
}
