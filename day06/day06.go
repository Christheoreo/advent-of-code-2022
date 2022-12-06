package day06

import (
	"strings"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

func SolveFirst(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	line := data[0]

	for i := 0; i < len(line); i++ {
		subset := line[i : i+4]

		bad := false
		for _, val := range subset {
			count := strings.Count(subset, string(val))
			if count > 1 {
				bad = true
				break
			}
		}
		if bad {
			continue
		}
		return i + 4
	}

	return 0
}
func SolveSecond(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	line := data[0]

	for i := 0; i < len(line); i++ {
		subset := line[i : i+14]
		bad := false
		for _, val := range subset {
			count := strings.Count(subset, string(val))
			if count > 1 {
				bad = true
				break
			}
		}
		if bad {
			continue
		}
		return i + 14
	}

	return 0
}
