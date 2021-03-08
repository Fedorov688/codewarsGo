package toNato

func toNato(words string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett",
		"Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor",
		"Whiskey", "X-ray", "Yankee", "Zulu"}
	var result string
	for key, word := range words {
		if 65 <= word && word <= 90 {
			result += nato[word-65]
		} else if 97 <= word && word <= 122 {
			result += nato[word-97]
		} else {
			if string(word) != " " {
				result += string(word)
			}
		}
		if key != len(words)-1 && string(word) != " " {
			result += " "
		}
	}
	return result
}
