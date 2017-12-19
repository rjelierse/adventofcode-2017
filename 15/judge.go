package day15

const startA = 65
const multiplierA = 16807
const dividerA = 4

const startB = 8921
const multiplierB = 48271
const dividerB = 8

const sampleSize = 40 * 1000 * 1000

// Mask for the lowest 16 bits
const mask = (1 << 16) - 1

type Judge struct {
	GenA *Generator
	GenB *Generator
}

func NewJudge(startA int, startB int) *Judge {
	judge := &Judge{}
	judge.GenA = NewGenerator(startA, multiplierA, dividerA)
	judge.GenB = NewGenerator(startB, multiplierB, dividerB)
	return judge
}

func (judge *Judge) FindMatches(sample int, shouldDivide bool) (matches int) {
	for i := 0; i < sample; i++ {
		if judge.round(shouldDivide) {
			matches++
		}
	}
	return
}

func (judge *Judge) round(shouldDivide bool) bool {
	a := judge.GenA.Next(shouldDivide)
	b := judge.GenB.Next(shouldDivide)
	return (a & mask) == (b & mask)
}
