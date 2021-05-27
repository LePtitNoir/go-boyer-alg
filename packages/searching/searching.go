package searching

type SearchAlgorithm interface {
	Indexes(text string, need string) []int
}

type SearchingContext struct {
	algorithm SearchAlgorithm
}

func (s *SearchingContext) SetAlgorithm(alg SearchAlgorithm) {
	s.algorithm = alg
}

func (s *SearchingContext) Match(text string, need string) bool {
	return len(s.algorithm.Indexes(text, need)) >= 0
}

func (s *SearchingContext) FindString(text string, need string, occ int) int {
	if tab := s.algorithm.Indexes(text, need); len(tab) >= occ {
		return tab[occ-1]
	}
	return -1
}

func (s *SearchingContext) FindAllString(text string, need string) []int {
	return s.algorithm.Indexes(text, need)
}
