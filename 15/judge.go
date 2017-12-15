package generator

const startA = 65
const multiplierA = 16807

const startB = 8921
const multiplierB = 48271

const sampleSize = 40 * 1000 * 1000

// Mask for the lowest 16 bits
const mask = (1 << 16) - 1

type Judge struct {
	GenA *Generator
	GenB *Generator
}

func NewJudge(startA int, startB int) *Judge {
	judge := &Judge{}
	judge.GenA = NewGenerator(startA, multiplierA)
	judge.GenB = NewGenerator(startB, multiplierB)
	return judge
}

func (judge *Judge) FindMatches(sample int) (matches int) {
	for i := 0; i < sample; i++ {
		if judge.round() {
			matches++
		}
	}
	return
}

func (judge *Judge) round() bool {
	a := judge.GenA.Next()
	b := judge.GenB.Next()
	return (a & mask) == (b & mask)
}
