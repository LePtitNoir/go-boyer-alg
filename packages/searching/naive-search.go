package searching

type NaiveSearch struct {
}

func (s *NaiveSearch) Indexes(text string, need string) []int {
	var tab []int
	indexWord := 0

	for indexText, letter := range text {
		if byte(letter) == need[indexWord] {
			indexWord++
			if len(need) == indexWord {
				indexWord = 0
				tab = append(tab, indexText-1)
			}
		} else {
			indexWord = 0
		}
	}

	return tab
}
