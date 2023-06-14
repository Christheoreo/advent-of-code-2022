package day03

import (
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func SolveFirst(data string) int {
	rucksacks := strings.Split(data, "\n")
	answer := 0

	for _, value := range rucksacks {
		halfwayPoint := len(value) / 2

		a := value[:halfwayPoint]
		b := value[halfwayPoint:]

		for _, v := range b {
			if strings.ContainsRune(a, v) {
				answer += strings.IndexRune(alphabet, v) + 1
				break
			}
		}
	}
	return answer
}

func SolveSecond(data string) int {
	rucksacks := strings.Split(data, "\n")
	answer := 0

	for i := 0; i < len(rucksacks); i += 3 {
		a, b, c := rucksacks[i], rucksacks[i+1], rucksacks[i+2]
		for _, r := range a {
			if strings.ContainsRune(b, r) && strings.ContainsRune(c, r) {
				answer += strings.IndexRune(alphabet, r) + 1
				break
			}
		}
	}
	return answer
}
