package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

type Stack struct {
	Containers []string
}

func (s *Stack) Move(amount int, to *Stack, asUnit bool) {
	containersLength := len(s.Containers)
	containersToMove := s.Containers[containersLength-amount:]
	s.Containers = s.Containers[0 : containersLength-amount]
	if asUnit {
		to.Containers = append(to.Containers, containersToMove...)
	} else {

		for i := len(containersToMove) - 1; i >= 0; i-- {
			to.Containers = append(to.Containers, containersToMove[i])
		}
	}
}

func SolveFirst(filename string) string {
	return helper(filename, true)
}
func SolveSecond(filename string) string {
	return helper(filename, false)
}

func helper(filename string, isFirstProb bool) string {
	stacks := map[int]*Stack{}
	data, _ := filereader.ReadFileToStringArray(filename)
	answer := ""
	maxKey := 1
	for lineIndex, val := range data {
		// Checking if this is the line that containers the stack numbers
		_, err := strconv.Atoi(string(val[1]))

		// Not a stack if there's an error
		if err != nil {
			continue
		}

		// this only works for stacks less than 10 - bad coding really
		for i, v := range val {
			num, err := strconv.Atoi(string(v))

			if err != nil {
				continue
			}

			// set the max key - looping through map gives inconsistant orders.
			// This way we know stacks start at 1 and end at X
			maxKey = num

			s := &Stack{
				Containers: make([]string, 0),
			}

			// loop backwards and populate the stack
			for ii := lineIndex - 1; ii >= 0; ii-- {
				str := string(data[ii][i])

				if str != "" && str != " " {
					s.Containers = append(s.Containers, str)
				}
			}

			stacks[num] = s
		}

		// getting the instruction data
		for i := lineIndex + 2; i < len(data); i++ {
			line := data[i]
			line = strings.Replace(line, "move", "", 1)
			line = strings.Replace(line, "from", ",", 1)
			line = strings.Replace(line, "to", ",", 1)
			line = strings.Replace(line, " ", "", -1)

			arr := strings.Split(line, ",")

			num1, _ := strconv.Atoi(arr[0])
			num2, _ := strconv.Atoi(arr[1])
			num3, _ := strconv.Atoi(arr[2])

			// we only want to "moveAsUnit" or the second problem.
			asUnit := !isFirstProb
			stacks[num2].Move(num1, stacks[num3], asUnit)
		}
		break
	}
	// loop through the stacks IN ORDER and get the top values.
	for i := 1; i <= maxKey; i++ {
		answer = fmt.Sprintf("%s%s", answer, stacks[i].Containers[len(stacks[i].Containers)-1])
	}
	return answer
}
