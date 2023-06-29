package day06

func mainLogic(data string, charsLength int) int {
	for i := 0; i <= len(data)-charsLength; i++ {
		chars := data[i : i+charsLength]
		match := true
		for x := 0; x < len(chars); x++ {
			if !match {
				break
			}
			for y := x + 1; y < len(chars); y++ {
				if !match {
					break
				}
				// maybe could be sped up if we know where the collision happened and increment to that bit +1?
				// if clast happends on 13 and 14, we know what can then start at 14 etc
				if chars[x] == chars[y] {
					match = false
					break
				}
			}
		}

		if match {
			return i + charsLength
		}

	}
	return 0
}

// much slower
// func SolveFirst(data string) int {
// 	return mainLogic(data, 4)
// }

func SolveFirst(data string) int {
	for i := 0; i <= len(data)-4; i++ {
		chars := data[i : i+4]
		if chars[0] != chars[1] &&
			chars[0] != chars[2] &&
			chars[0] != chars[3] &&
			chars[1] != chars[2] &&
			chars[1] != chars[3] &&
			chars[2] != chars[3] {
			return i + 4
		}
	}
	return 0
}
func SolveSecond(data string) int {
	return mainLogic(data, 14)
}

// Made a script to write the if statement lol - was curious about speeds
func SolveSecondHardCoded(data string) int {
	for i := 0; i <= len(data)-14; i++ {
		chars := data[i : i+14]
		if chars[0] != chars[1] &&
			chars[0] != chars[2] &&
			chars[0] != chars[3] &&
			chars[0] != chars[4] &&
			chars[0] != chars[5] &&
			chars[0] != chars[6] &&
			chars[0] != chars[7] &&
			chars[0] != chars[8] &&
			chars[0] != chars[9] &&
			chars[0] != chars[10] &&
			chars[0] != chars[11] &&
			chars[0] != chars[12] &&
			chars[0] != chars[13] &&
			chars[1] != chars[2] &&
			chars[1] != chars[3] &&
			chars[1] != chars[4] &&
			chars[1] != chars[5] &&
			chars[1] != chars[6] &&
			chars[1] != chars[7] &&
			chars[1] != chars[8] &&
			chars[1] != chars[9] &&
			chars[1] != chars[10] &&
			chars[1] != chars[11] &&
			chars[1] != chars[12] &&
			chars[1] != chars[13] &&
			chars[2] != chars[3] &&
			chars[2] != chars[4] &&
			chars[2] != chars[5] &&
			chars[2] != chars[6] &&
			chars[2] != chars[7] &&
			chars[2] != chars[8] &&
			chars[2] != chars[9] &&
			chars[2] != chars[10] &&
			chars[2] != chars[11] &&
			chars[2] != chars[12] &&
			chars[2] != chars[13] &&
			chars[3] != chars[4] &&
			chars[3] != chars[5] &&
			chars[3] != chars[6] &&
			chars[3] != chars[7] &&
			chars[3] != chars[8] &&
			chars[3] != chars[9] &&
			chars[3] != chars[10] &&
			chars[3] != chars[11] &&
			chars[3] != chars[12] &&
			chars[3] != chars[13] &&
			chars[4] != chars[5] &&
			chars[4] != chars[6] &&
			chars[4] != chars[7] &&
			chars[4] != chars[8] &&
			chars[4] != chars[9] &&
			chars[4] != chars[10] &&
			chars[4] != chars[11] &&
			chars[4] != chars[12] &&
			chars[4] != chars[13] &&
			chars[5] != chars[6] &&
			chars[5] != chars[7] &&
			chars[5] != chars[8] &&
			chars[5] != chars[9] &&
			chars[5] != chars[10] &&
			chars[5] != chars[11] &&
			chars[5] != chars[12] &&
			chars[5] != chars[13] &&
			chars[6] != chars[7] &&
			chars[6] != chars[8] &&
			chars[6] != chars[9] &&
			chars[6] != chars[10] &&
			chars[6] != chars[11] &&
			chars[6] != chars[12] &&
			chars[6] != chars[13] &&
			chars[7] != chars[8] &&
			chars[7] != chars[9] &&
			chars[7] != chars[10] &&
			chars[7] != chars[11] &&
			chars[7] != chars[12] &&
			chars[7] != chars[13] &&
			chars[8] != chars[9] &&
			chars[8] != chars[10] &&
			chars[8] != chars[11] &&
			chars[8] != chars[12] &&
			chars[8] != chars[13] &&
			chars[9] != chars[10] &&
			chars[9] != chars[11] &&
			chars[9] != chars[12] &&
			chars[9] != chars[13] &&
			chars[10] != chars[11] &&
			chars[10] != chars[12] &&
			chars[10] != chars[13] &&
			chars[11] != chars[12] &&
			chars[11] != chars[13] &&
			chars[12] != chars[13] {
			return i + 14
		}
	}
	return 0
}
