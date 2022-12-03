package day03

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func SolveFirst(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	points := 0
	for _, v := range data {
		containerOne := v[:len(v)/2]
		containerTwo := v[len(v)/2:]

		str := fmt.Sprintf("[%s]", strings.Join(strings.Split(containerOne, ""), ","))

		// Skipping error check as i know this works
		re, _ := regexp.Compile(str)

		foundLetter := string(re.Find([]byte(containerTwo)))

		index := strings.Index(alphabet, foundLetter)

		points += index + 1
	}
	return points
}

// theres definitely a better way to do this
func SolveSecond(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	points := 0

	threes := make([][3]string, 1)
	currentI := 0
	currentII := 0

	for _, v := range data {
		threes[currentI][currentII] = v
		if currentII == 2 {
			currentI++
			currentII = 0
			threes = append(threes, [3]string{"", "", ""})
		} else {
			currentII++
		}
	}

	for _, v := range threes {
		for _, character := range v[0] {
			if strings.ContainsRune(v[1], character) && strings.ContainsRune(v[2], character) {
				points += strings.Index(alphabet, string(character)) + 1
				break
			}
		}
	}
	return points
}
