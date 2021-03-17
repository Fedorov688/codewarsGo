package playingWithDigits

import (
	"math"
	"strconv"
)

func playingWithDigits(n, p int) int {

	var numIntList = strconv.Itoa(n)
	var tmp int
	var result int

	for _, value := range numIntList {
		tmp, _ = strconv.Atoi(string(value))
		result += int(math.Pow(float64(tmp), float64(p)))
		p++
	}

	if result%n == 0 {
		return result / n
	} else {
		return -1
	}
}
