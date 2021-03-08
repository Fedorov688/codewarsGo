package findUniq

func FindUniq(arr []float32) float32 {
	var resultArr = map[float32]int{}

	if len(arr) > 1 {
		for _, value := range arr {
			if v, ok := resultArr[value]; ok {
				resultArr[value] = v + 1
			} else {
				resultArr[value] = 1
			}
		}
		for key, value := range resultArr {
			if value == 1 {
				return key
			}
		}
	} else {
		return arr[0]
	}

	return arr[0]
}
