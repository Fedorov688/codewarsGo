package createPhoneNumber

import (
	"fmt"
)

func arrToStr(arr []uint) string {
	var tmp string
	for _, val := range arr {
		tmp += fmt.Sprintf("%d", val)
	}
	return tmp
}

func createPhoneNumber(numbers [10]uint) string {
	return fmt.Sprintf("(%s) %s-%s",
		arrToStr(numbers[0:3]),
		arrToStr(numbers[3:6]),
		arrToStr(numbers[6:10]))
}
