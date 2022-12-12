package day11

import (
	"fmt"

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

func SolveSecond(filename string, turns int) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	monkeys := setupMonkeysA(data)

	for turn := 0; turn < turns; turn++ {

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
