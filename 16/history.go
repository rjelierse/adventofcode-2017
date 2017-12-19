package day16

type History struct {
	positions []string
}

func NewHistory() *History {
	h := new(History)
	h.positions = make([]string, 0)
	return h
}

func (h *History) Save(position string) {
	h.positions = append(h.positions, position)
}

func (h *History) SeenBefore(position string) (bool, int) {
	for index, p := range h.positions {
		if p == position {
			return true, index
		}
	}

	return false, -1
}

func (h *History) Get(index int) string {
	return h.positions[index]
}
