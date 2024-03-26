package decode

import (
	"strings"
)

func AlienDecode(input string) int {

	if input == "" {
		return 0
	}

	deocodeTemplate := map[string]int{"A": 1, "B": 5, "Z": 10, "L": 50, "C": 100, "D": 500, "R": 1000}
	var result int
	inputSlice := strings.Split(input, "")

	for i := 0; i < len(inputSlice); i++ {
		if deocodeTemplate[inputSlice[i]] == 0 {
			return 0
		}

		if i+1 < len(inputSlice) && deocodeTemplate[inputSlice[i]] < deocodeTemplate[inputSlice[i+1]] {
			result += deocodeTemplate[inputSlice[i+1]] - deocodeTemplate[inputSlice[i]]
			i++
		} else {

			result += deocodeTemplate[inputSlice[i]]
		}
	}
	return result
}
