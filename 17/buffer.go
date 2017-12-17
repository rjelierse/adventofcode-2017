package spinlock

type Ring struct {
	buffer   []int
	value    int
	position int
}

func NewRing() *Ring {
	r := new(Ring)
	r.buffer = []int{0}
	r.value = 0
	r.position = 0
	return r
}

func (r *Ring) Run(size int, steps int) {
	for r.value < steps {
		r.Step(size)
	}
}

func (r *Ring) Step(size int) {
	position := r.nextPosition(size)
	value := r.nextValue()
	r.insertAt(value, position)
	r.value = value
	r.position = position
}

func (r *Ring) ValueAfter(value int) int {
	for i, v := range r.buffer {
		if value == v {
			return r.buffer[i+1]
		}
	}
	return -1
}

func (r *Ring) nextPosition(length int) int {
	return ((r.position + length) % len(r.buffer)) + 1
}

func (r *Ring) nextValue() int {
	return r.value + 1
}

func (r *Ring) insertAt(value int, position int) {
	if position == len(r.buffer) {
		r.buffer = append(r.buffer, value)
	} else {
		r.buffer = append(r.buffer[:position], append([]int{value}, r.buffer[position:]...)...)
	}
}
