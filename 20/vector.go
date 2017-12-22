package day20

type Vector struct {
	X int
	Y int
	Z int
}

// Sub takes the components from the vector and subtracts using the components from the other vector
func (v Vector) Sub(b Vector) Vector {
	return Vector{v.X - b.X, v.Y - b.Y, v.Z - b.Z}
}

// Sum takes the components for the vector and calculates the absolute distance to <0,0,0>
func (v Vector) Sum() int {
	abs := func(v int) int {
		if v >= 0 {
			return v
		}
		return v * -1
	}

	return abs(v.X) + abs(v.Y) + abs(v.Z)
}
