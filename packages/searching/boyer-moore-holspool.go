package searching

type BoyerMooreHorspoolSearch struct {
}

func (s *BoyerMooreHorspoolSearch) Indexes(text string, key string) []int {

	if len(key) > len(text) {
		return nil
	}

	var positions []int
	iText, iKey := len(key)-1, len(key)-1
	tabDump := make(map[string]int)

	for i, l := range key {
		tabDump[string(l)] = len(key) - (i + 1)
	}

	for iText < len(text) {
		if text[iText] == key[iKey] {
			if iKey == 0 {
				positions = append(positions, iText)
				iKey = len(key) - 1
				iText = iText + len(key)
			} else {
				iText--
				iKey--
			}
		} else {
			iKey = len(key) - 1

			if v, ok := tabDump[string(text[iText])]; ok {
				iText = iText + v
			} else {
				iText = iText + len(key)
			}
		}
	}

	return positions
}
