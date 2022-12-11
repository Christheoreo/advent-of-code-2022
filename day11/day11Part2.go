package day11

import (
	"fmt"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

func SolveSecond(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	monkeys := setupMonkeys(data)

	for turn := 0; turn < 1000; turn++ {

		for _, monkey := range monkeys {
			monkey.InspectionCount += len(monkey.StartingItems)

			for _, item := range monkey.StartingItems {
				var newItem int
				if monkey.Operation == "*" {
					if monkey.OperationNumber == -1 {
						// fmt.Printf("Worry level is multiplied by %d to %d\n", item, item*item)
						fmt.Printf("%d and %d = %d\n", item, item, item*item)
						newItem = item * item
					} else {
						// fmt.Printf("Worry level is multiplied by %d to %d\n", item, item*monkey.OperationNumber)
						fmt.Printf("%d and %d = %d\n", item, monkey.OperationNumber, item*monkey.OperationNumber)
						newItem = item * monkey.OperationNumber
					}
				} else if monkey.Operation == "+" {
					if monkey.OperationNumber == -1 {
						// fmt.Printf("Worry level increases by %d to %d\n", item, item+item)
						newItem = item + item
					} else {
						// fmt.Printf("Worry level increases by %d to %d\n", item, item+monkey.OperationNumber)
						newItem = item + monkey.OperationNumber
					}
				}
				// now test it

				// alway is diviisble

				if newItem%monkey.TestNumber == 0 {
					// fmt.Printf("Current worry level is divisible by %d.\n", monkey.TestNumber)
					// fmt.Printf("Item with worry level %d is thrown to monkey %d\n", item, monkey.TrueOutcome)
					// pass the item to another monkey
					monkeys[monkey.TrueOutcome].StartingItems = append(monkeys[monkey.TrueOutcome].StartingItems, newItem)
				} else {
					// fmt.Printf("Current worry level is not divisible by %d.\n", monkey.TestNumber)
					// fmt.Printf("Item with worry level %d is thrown to monkey %d\n", item, monkey.FalseOutcome)
					monkeys[monkey.FalseOutcome].StartingItems = append(monkeys[monkey.FalseOutcome].StartingItems, newItem)
				}
			}
			// clear this monkeys items
			monkey.StartingItems = make([]int, 0)
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
	return "1"
}

func addStrings(stringOne string, stringTwo string) string {
	return "0"
}

func isDivisible(stringOne string, stringTwo string) bool {
	return true
}
