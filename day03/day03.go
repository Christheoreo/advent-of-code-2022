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
		// fmt.Println(v)

		maxLength := len(v[0])
		maxArr := v[0]

		if len(v[1]) > maxLength {
			maxLength = len(v[1])
			maxArr = v[1]
		}

		if len(v[2]) > maxLength {
			maxLength = len(v[2])
			maxArr = v[2]
		}

		for i := 0; i < maxLength; i++ {
			if strings.Contains(v[1], string(maxArr[i])) && strings.Contains(v[2], string(maxArr[i])) {
				fmt.Println(string(maxArr[i]))

				points += strings.Index(alphabet, string(maxArr[i])) + 1
				break
			}

			// fmt.Println(string(str[i]))

			// points += strings.Index(alphabet, string(str[i])) + 1
			break
		}
	}

	for i, _ := range data {
		// fmt.Println(i + 1)
		// fmt.Println((i + 1) % 3)
		if (i+1)%3 != 0 {
			// fmt.Println("ergh")
			continue
		}

		// str := fmt.Sprintf("%s%s%s", data[i], data[i-1], data[i-2])
		a, b, c := data[i], data[i-1], data[i-2]
		fmt.Printf("%s\n%s\n%s", a, b, c)

		for k := 0; k < len(a); k++ {
			if strings.Contains(b, string(a[k])) && strings.Contains(c, string(a[k])) {
				fmt.Println(string(a[k]))

				points += strings.Index(alphabet, string(a[k])) + 1
				break
			}

			// fmt.Println(string(str[i]))

			// points += strings.Index(alphabet, string(str[i])) + 1
			break
		}

		// fmt.Println(str)

		// for i := 0; i < len(str); i++ {
		// 	count := strings.Count(str, string(str[i]))

		// 	if count != 3 {
		// 		continue
		// 	}

		// 	fmt.Println(string(str[i]))

		// 	points += strings.Index(alphabet, string(str[i])) + 1
		// 	break
		// }
	}
	return points
}
