package day03

import (
	"math"
)

const (
	Bottom = iota
	Left
	Top
	Right
)

func position(value int) (x int, y int) {
	if value == 1 {
		return 0, 0
	}

	base := int(math.Ceil(math.Sqrt(float64(value))))

	if base%2 == 0 {
		base = base + 1
	}

	maxValue := base * base
	maxCoord := (base - 1) / 2

	if value == maxValue {
		return maxCoord, maxCoord
	}

	if value == maxValue+1 {
		return maxCoord + 1, maxCoord
	}

	offset := maxValue - value

	switch offset / base {
	case Bottom:
		x = maxCoord - (offset % base)
		y = maxCoord
	case Left:
		x = -maxCoord
		y = maxCoord - (offset % base)
	case Top:
		x = (offset % base) - maxCoord
		y = -maxCoord
	case Right:
		x = maxCoord
		y = (offset % base) - maxCoord
	}

	return
}
