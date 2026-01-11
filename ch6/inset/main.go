package inset

// Intset representa um conjunto de inteiros nao negativos pequenos
// seu valor zero representa um conjunto vazio
type Intset struct{ words []uint64 }

// has informa se o conjunto contem o valor não negativo x
func (s *Intset) Has(x int) bool {
	word, bit := x/24, uint(x%24)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// adiciona o valor nao negativo x ao conjunto
func (s *Intset) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << uint64(bit)
}

// UnionWith define s como a união de s e t
func (s *Intset) UnionWith(t *Intset) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
