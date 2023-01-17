package day11

import (
	"strconv"
	"strings"
)

func sortData(data []string) []*Monkey {
	monkeys := make([]*Monkey, 0)
	var monkey *Monkey
	for index := 0; index < len(data); index++ {
		if !strings.Contains(data[index], "Monkey") {
			continue
		}
		if index != 0 {
			monkeys = append(monkeys, monkey)
		}
		monkey = &Monkey{
			Name:            data[index][7 : len(data[index])-1],
			WorryLevels:     make([]int, 0),
			MonkeyToThrowTo: [2]int{0, 0},
		}
		// Index + 1 will be the worry level
		worryLevel := strings.Split(data[index+1][18:], ", ")
		for _, wLS := range worryLevel {
			val, _ := strconv.Atoi(wLS)
			monkey.WorryLevels = append(monkey.WorryLevels, val)
		}

		// Index + 2 will be the operation
		operation := data[index+2][13:]
		if operation == "new = old * 19" {
			monkey.Operation = 'a'
		} else if operation == "new = old + 6" {
			monkey.Operation = 'b'
		} else if operation == "new = old * old" {
			monkey.Operation = 'c'
		} else {
			monkey.Operation = 'd'
		}

		// Index + 3 will be the Test Operation
		val, _ := strconv.Atoi(data[index+3][21:])
		monkey.TestOperation = val

		// Index + 4 will be the true art of the test
		val, _ = strconv.Atoi(data[index+4][29:])
		monkey.MonkeyToThrowTo[0] = val

		// Index + 5 will be the false part of the Test
		val, _ = strconv.Atoi(data[index+5][30:])
		monkey.MonkeyToThrowTo[1] = val
	}
	// Then we need to append the monkey for the last one.
	monkeys = append(monkeys, monkey)
	return monkeys
}

// func sortData(data []string) []*Monkey {
// 	monkeys := make([]*Monkey, 0)
// 	var monkey *Monkey
// 	for index, line := range data {
// 		if !strings.Contains(line, "Monkey") {
// 			continue
// 		}
// 		if index != 0 {
// 			monkeys = append(monkeys, monkey)
// 		}
// 		monkey = &Monkey{
// 			Name:            line[7 : len(line)-1],
// 			WorryLevels:     make([]int, 0),
// 			MonkeyToThrowTo: [2]int{0, 0},
// 		}
// 		// Index + 1 will be the worry level
// 		origonalWorryLevels := strings.Split(data[index+1][18:], ", ")

// 		for _, wLS := range origonalWorryLevels {
// 			val, _ := strconv.Atoi(wLS)
// 			monkey.WorryLevels = append(monkey.WorryLevels, val)
// 		}

// 		operation := data[index+2][13:]
// 		if operation == "new = old * 19" {
// 			monkey.Operation = 'a'
// 		} else if operation == "new = old + 6" {
// 			monkey.Operation = 'b'
// 		} else if operation == "new = old * old" {
// 			monkey.Operation = 'c'
// 		} else {
// 			monkey.Operation = 'd'
// 		}

// 		val, _ := strconv.Atoi(data[index+3][21:])
// 		monkey.TestOperation = val

// 		//29 - truw
// 		val, _ = strconv.Atoi(data[index+4][29:])
// 		monkey.MonkeyToThrowTo[0] = val
// 		//30 - false
// 		val, _ = strconv.Atoi(data[index+5][30:])
// 		monkey.MonkeyToThrowTo[1] = val

// 	}
// 	monkeys = append(monkeys, monkey)
// 	return monkeys
// }
