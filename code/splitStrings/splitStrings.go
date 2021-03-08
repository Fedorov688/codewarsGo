package splitStrings

func Solution(str string) []string {
	var lenght = len(str)
	var resultSlice []string
	if lenght%2 == 0 {
		for key := 0; key < lenght; key += 2 {
			resultSlice = append(resultSlice, str[key:key+2])
		}
	} else {
		for key := 0; key < lenght-1; key += 2 {
			resultSlice = append(resultSlice, str[key:key+2])
		}
		resultSlice = append(resultSlice, str[lenght-1:lenght]+"_")
	}
	return resultSlice
}
