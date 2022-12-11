package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

type MonkeyA struct {
	StartingItems   []string
	Operation       string
	OperationNumber string
	Test            string
	TestNumber      string
	TrueOutcome     int
	FalseOutcome    int
	InspectionCount int
}

func SolveSecond(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	monkeys := setupMonkeysA(data)

	for turn := 0; turn < 10000; turn++ {

		fmt.Println(turn)

		for _, monkey := range monkeys {
			monkey.InspectionCount += len(monkey.StartingItems)

			for _, item := range monkey.StartingItems {
				var newItem string
				if monkey.Operation == "*" {
					if monkey.OperationNumber == "-1" {
						// fmt.Printf("Worry level is multiplied by %d to %d\n", item, item*item)
						newItem = multiplyStrings(item, item)
						// fmt.Printf("%s and %s = %s\n", item, item, newItem)
					} else {
						// fmt.Printf("Worry level is multiplied by %d to %d\n", item, item*monkey.OperationNumber)
						newItem = multiplyStrings(item, monkey.OperationNumber)
						// fmt.Printf("%s and %s = %s\n", item, monkey.OperationNumber, newItem)
					}
				} else if monkey.Operation == "+" {
					if monkey.OperationNumber == "-1" {
						newItem = addStrings(item, item)
						// fmt.Printf("Worry level increases by %s to %s\n", item, newItem)
					} else {
						newItem = addStrings(item, monkey.OperationNumber)
						// fmt.Printf("Worry level increases by %s to %s\n", item, newItem)
					}
				}
				// now test it

				// alway is diviisble
				// fmt.Printf("dividing\n")
				if isDivisible(newItem, monkey.TestNumber) {
					// fmt.Printf("Current worry level is divisible by %d.\n", monkey.TestNumber)
					// fmt.Printf("Item with worry level %d is thrown to monkey %d\n", item, monkey.TrueOutcome)
					// pass the item to another monkey
					monkeys[monkey.TrueOutcome].StartingItems = append(monkeys[monkey.TrueOutcome].StartingItems, newItem)
				} else {
					// fmt.Printf("Current worry level is not divisible by %d.\n", monkey.TestNumber)
					// fmt.Printf("Item with worry level %d is thrown to monkey %d\n", item, monkey.FalseOutcome)
					monkeys[monkey.FalseOutcome].StartingItems = append(monkeys[monkey.FalseOutcome].StartingItems, newItem)
				}
				// fmt.Printf("Done dividing\n")
			}
			// clear this monkeys items
			monkey.StartingItems = make([]string, 0)
		}
	}

	a, b := 0, 0

	for i, m := range monkeys {
		fmt.Printf("monkey %d inspected items %d times.\n", i, m.InspectionCount)
		if m.InspectionCount > a {
			b = a
			a = m.InspectionCount
			continue
		}
		if m.InspectionCount >= b {
			b = m.InspectionCount
		}
	}
	return a * b
}

func multiplyStrings(stringOne string, stringTwo string) string {

	// we need to start from the smallest
	smallest := ""
	largest := ""

	if len(stringOne) < len(stringTwo) {
		smallest = stringOne
		largest = stringTwo
	} else {
		smallest = stringTwo
		largest = stringOne
	}
	answer := ""

	for s, sI := len(smallest)-1, 0; s >= 0; s, sI = s-1, sI+1 {
		sNum, _ := strconv.Atoi(string(smallest[s]))
		for l, lI := len(largest)-1, 0; l >= 0; l, lI = l-1, lI+1 {
			lNum, _ := strconv.Atoi(string(largest[l]))

			sum := sNum * lNum

			sumString := fmt.Sprint(sum)
			for i := 0; i < sI+lI; i++ {
				sumString += "0"
			}

			answer = addStrings(answer, sumString)

		}
	}
	return answer
}

func addStrings(stringOne string, stringTwo string) string {
	for len(stringTwo) < len(stringOne) {
		stringTwo = "0" + stringTwo
	}
	for len(stringOne) < len(stringTwo) {
		stringOne = "0" + stringOne
	}

	carry := 0
	answer := ""
	for i := len(stringOne) - 1; i >= 0; i-- {
		numOne, _ := strconv.Atoi(string(stringOne[i]))
		numTwo, _ := strconv.Atoi(string(stringTwo[i]))

		sum := numOne + numTwo + carry
		if sum > 9 {
			carry = 1
			sum = sum - 10
		} else {
			carry = 0
		}
		answer = fmt.Sprintf("%d%s", sum, answer)
	}

	if carry > 0 {
		answer = fmt.Sprintf("%d%s", carry, answer)
	}

	return answer
}

func subTractStrings(stringOne string, stringTwo string) string {
	for len(stringTwo) < len(stringOne) {
		stringTwo = "0" + stringTwo
	}
	for len(stringOne) < len(stringTwo) {
		stringOne = "0" + stringOne
	}

	carry := 0
	answer := ""
	for i := len(stringOne) - 1; i >= 0; i-- {
		numOne, _ := strconv.Atoi(string(stringOne[i]))
		numTwo, _ := strconv.Atoi(string(stringTwo[i]))
		numTwo += carry

		if numTwo == 10 {
			carry = 1
			answer = fmt.Sprintf("%d%s", numOne, answer)
			continue
		}
		// if numTwo > numOne {
		// 	numOne, _ = strconv.Atoi(string(s0 - sumtringOne[i : i+1]))
		// }

		sum := numOne - numTwo

		if sum < 0 {
			carry = 1
			// numOne, _ = strconv.Atoi(string(stringOne[i : i+1]))
			// sum = numOne - numTwo
			// here
			// sum = sum - 10
			n := 10 + sum
			answer = fmt.Sprintf("%d%s", n, answer)
		} else {
			carry = 0
			answer = fmt.Sprintf("%d%s", sum, answer)
		}
	}

	// if carry > 0 {
	// 	answer = fmt.Sprintf("%d%s", carry, answer)
	// }

	if string(answer[0]) == "0" {
		if len(answer) == 1 {
			return "0"
		}
		return answer[1:]
	}

	return answer
}

// this needs tweaking - better way to get this done
func isDivisible(stringOne string, by string) bool {
	// i := stringOne
	// fmt.Println(stringOne)

	// NOTE - try and do {by} with as many zeros as possible, until we get down to as smaller as possible.
	x := by
	for i := 0; i < len(stringOne)-3; i++ {
		x += "0"
	}
	i := subTractStrings(stringOne, x)
	// fmt.Printf("i is now %s and was %s and is beoing divided by %s --- %s\n", i, stringOne, by, x)
	for {

		if len(i) < 7 {

			num, _ := strconv.Atoi(i)
			num2, _ := strconv.Atoi(by)

			return num%num2 == 0
		}

		length := len(i)

		b := by

		for k := 0; k < length-1-len(by); k++ {
			b += "0"
		}

		// b := by

		// for k := 0; k < len(i)-6; k++ {
		// 	b += "0"
		// }

		newI := subTractStrings(i, b)
		i = newI
		// fmt.Println(i)
	}
}

func setupMonkeysA(data []string) []*MonkeyA {
	monkeys := make([]*MonkeyA, 0)
	var monkey *MonkeyA
	for _, line := range data {
		if len(line) < 1 {
			// monkeys = append(monkeys, monkey)
			continue
		}
		if string(line[0]) == "M" {
			monkey = &MonkeyA{}
			monkey.StartingItems = make([]string, 0)
			continue
		}

		if strings.Contains(line, "items") {
			start := line[18:]

			numbersS := strings.Split(start, ", ")
			fmt.Println(numbersS)
			monkey.StartingItems = append(monkey.StartingItems, numbersS...)
		}

		if strings.Contains(line, "Operation") {
			thing := strings.Split(line, "=")[1]
			thing = strings.TrimSpace(thing)
			x := strings.Split(thing, " ")
			monkey.Operation = x[1]

			if x[2] == "old" {
				// use this as a special case
				monkey.OperationNumber = "-1"
			} else {
				monkey.OperationNumber = x[2]
			}
			continue
		}

		if strings.Contains(line, "Test") {
			//always seems rto be divisible
			monkey.TestNumber = line[21:]
			monkey.Test = "/"
			continue
		}

		if strings.Contains(line, "true") {
			num, _ := strconv.Atoi(line[29:])
			// fmt.Printf("true outcome = %d\n", num)
			monkey.TrueOutcome = num
			continue
		}

		if strings.Contains(line, "false") {
			num, _ := strconv.Atoi(line[30:])
			monkey.FalseOutcome = num
			// fmt.Printf("false outcome = %d\n", num)

			monkeys = append(monkeys, monkey)
			continue
		}

	}
	return monkeys
}
