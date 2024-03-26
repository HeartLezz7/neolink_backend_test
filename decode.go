package decode

import (
	"strings"
)

func AlienDecode(input string) int {
	template := map[string]int{"A": 1, "B": 5, "Z": 10, "L": 50, "C": 100, "D": 500, "R": 1000}
	var result int
	inputSlice := strings.Split(input, "")

	for i := 0; i < len(inputSlice); i++ {
		if i+1 < len(inputSlice) && template[inputSlice[i]] < template[inputSlice[i+1]] {
			result += template[inputSlice[i+1]] - template[inputSlice[i]]
			i++
		} else {

			result += template[inputSlice[i]]
		}
	}
	return result
}
