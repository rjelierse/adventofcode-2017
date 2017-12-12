package pipes

import (
	"strings"
	"strconv"
)

type Tree struct {
	programs map[int]*Program
}

type Program struct {
	pid   int
	pipes []int
}

func ParseGroup(line string) (*Program, error) {
	data := strings.Split(line, " <-> ")

	g := new(Program)

	v, err := strconv.ParseInt(data[0], 10, 0)
	if err != nil {
		return nil, err
	}
	g.pid = int(v)

	for _, l := range strings.Split(data[1], ", ") {
		v, err := strconv.ParseInt(l, 10, 0)
		if err != nil {
			return nil, err
		}

		g.pipes = append(g.pipes, int(v))
	}

	return g, nil
}

func ParseTree(lines []string) (*Tree, error) {
	tree := new(Tree)
	tree.programs = make(map[int]*Program)

	for _, line := range lines {
		g, err := ParseGroup(line)
		if err != nil {
			return nil, err
		}

		tree.programs[g.pid] = g
	}

	return tree, nil
}

func (t *Tree) NumGroups() int {
	allVisits := make([][]int, 0)
	for pid := range t.programs {
		needsVisit := true

		for _, group := range allVisits {
			if isVisited(pid, group) {
				needsVisit = false
				break
			}
		}

		if needsVisit {
			group := t.Group(pid)
			allVisits = append(allVisits, group)
		}
	}

	return len(allVisits)
}

func (t *Tree) Group(pid int) []int {
	return t.visit(t.programs[pid], make([]int, 0))
}

func (t *Tree) Size(pid int) int {
	return len(t.Group(pid))
}

func (t *Tree) visit(p *Program, visited []int) []int {
	if isVisited(p.pid, visited) {
		return visited
	}

	visited = append(visited, p.pid)

	for _, pid := range p.pipes {
		c := t.programs[pid]
		visited = t.visit(c, visited)
	}

	return visited
}

func isVisited(pid int, visited []int) bool {
	for _, cur := range visited {
		if cur == pid {
			return true
		}
	}

	return false
}
