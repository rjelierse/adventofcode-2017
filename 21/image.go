package day21

type Image Grid

const ImageStart = ".#./..#/###"

func (img Image) Count() (count int) {
	for _, row := range img {
		for _, char := range row {
			if char == '#' {
				count++
			}
		}
	}
	return
}

func (img Image) Enhance(list map[Pattern]Pattern) Image {
	width := len(img)
	// Determine size of squares
	var nGrids, sGrids int
	switch 0 {
	case width % 2:
		// Split into 2x2 parts
		nGrids = width / 2
		sGrids = 2
	case width % 3:
		// Split into 3x3 parts
		nGrids = width / 3
		sGrids = 3
	default:
		panic("image size not dividable through 2 or 3")
	}

	// Apply enhancements to square
	table := make([][]Grid, nGrids)
	for y := 0; y < nGrids; y++ {
		table[y] = make([]Grid, nGrids)
		for x := 0; x < nGrids; x++ {
			grid := NewGrid(sGrids)
			for j := 0; j < sGrids; j++ {
				for i := 0; i < sGrids; i++ {
					dx := x * sGrids + i
					dy := y * sGrids + j
					grid[j][i] = img[dy][dx]
				}
			}
			pattern := grid.ToPattern()
			replace, exists := list[pattern]
			if !exists {
				panic("pattern not found")
			}
			table[y][x] = replace.ToGrid()
		}
	}

	// Rebuild image from squares
	sGrids += 1
	img2 := NewGrid(nGrids * sGrids)
	for y, row := range table {
		for x, grid := range row {
			for j := 0; j < sGrids; j++ {
				for i := 0; i < sGrids; i++ {
					dx := x * sGrids + i
					dy := y * sGrids + j
					img2[dy][dx] = grid[j][i]
				}
			}
		}
	}
	return Image(img2)
}

func DefaultImage() Image {
	grid := Pattern(ImageStart).ToGrid()
	return Image(grid)
}
