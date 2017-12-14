package knot

type Knot struct {
	sparseHash      []int
	length          int
	currentPosition int
	skip            int
}

const cipherLength = 256
const hashingRounds = 64

func Hash(lengths []int) []int {
	lengths = append(lengths, 17, 31, 73, 47, 23)
	k := NewKnot(cipherLength)
	for r := 0; r < hashingRounds; r++ {
		k.Round(lengths)
	}
	return k.DenseHash()
}

func NewKnot(length int) *Knot {
	k := new(Knot)
	k.length = length
	k.sparseHash = make([]int, length)
	for i := 0; i < length; i++ {
		k.sparseHash[i] = i
	}

	return k
}

// Round runs a single round of transformations on the knot
func (k *Knot) Round(lengths []int) {
	for _, length := range lengths {
		k.Transform(length)
	}
}

// Transform takes a length of numbers from the sparseHash and applies them in reverse order
func (k *Knot) Transform(length int) {
	l := make([]int, length)

	// First, copy the values in reverse
	for i := 0; i < length; i++ {
		pos := (k.currentPosition + length - i - 1) % k.length
		l[i] = k.sparseHash[pos]
	}

	// Then, apply the values to the sparseHash
	for i := 0; i < length; i++ {
		pos := (k.currentPosition + i) % k.length
		k.sparseHash[pos] = l[i]
	}

	k.currentPosition = k.currentPosition + length + k.skip
	k.skip++
}

// Checksum calculates the checksum of the knot by taking the first two number of the sparseHash and multiplying them
func (k *Knot) Checksum() int {
	return k.sparseHash[0] * k.sparseHash[1]
}

const hashingGroupSize = 16

// DenseHash takes the sparse hash and transforms it by XORing each group of 16 bytes
func (k *Knot) DenseHash() []int {
	hash := make([]int, k.length / hashingGroupSize)

	for i := 0; i < hashingGroupSize; i++ {
		start := i * hashingGroupSize
		end := start + hashingGroupSize

		hash[i] = XorAll(k.sparseHash[start:end])
	}

	return hash
}

// XorAll takes all values in the slice and XOR's them in order.
func XorAll(values []int) int {
	result := values[0]

	for i := 1; i < len(values); i++ {
		result = result ^ values[i]
	}

	return result
}
