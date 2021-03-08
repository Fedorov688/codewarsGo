package rgbToHex

import "fmt"

func rgb(r int, g int, b int) string {
	var result string
	rgbArray := []int{r, g, b}
	for _, value := range rgbArray {
		if value <= 0 {
			result += "00"
		} else if value >= 255 {
			result += "FF"
		} else {
			result += fmt.Sprintf("%X", value)
		}
	}
	return result
}
