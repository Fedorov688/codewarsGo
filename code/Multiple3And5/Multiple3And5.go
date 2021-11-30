package Multiple3And5

func Multiple3And5(number int) (result int) {
	if number < 3 {
		return
	}

	for i := 3; i < number; i++ {
		if i%3 == 0 {
			result += i
		} else if i%5 == 0 {
			result += i
		}
	}
	return
}
