package sumOfIntInStr

import "strconv"

func SumOfIntegersInString(strng string) int {
	var tmp string
	var res int
	for _, val := range strng {
		if 48 <= val && val <= 57 {
			tmp += string(val)
		} else {
			t, _ := strconv.Atoi(tmp)
			res += t
			tmp = ""
		}
	}

	if tmp != "" {
		t, _ := strconv.Atoi(tmp)
		res += t
		tmp = ""
	}

	return res
}
