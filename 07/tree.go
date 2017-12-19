package day07

// BuildTree constructs the tree of programs and return the root node
func BuildTree(lines []string) *Program {
	m := make(map[string]*Program)
	c := make([]*Program, 0)

	// Populate map of programs
	for _, line := range lines {
		p, err := ParseLine(line)
		if err != nil {
			continue
		}

		m[p.Name] = p

		if len(p.Children) > 0 {
			c = append(c, p)
		}
	}

	// Iterate programs to build tree
	for _, p := range m {
		for _, c := range p.Children {
			p.AddChild(m[c])
		}
	}

	// Iterate all nodes that have children, and return node that does not have a parent
	for _, p := range c {
		if p.Parent == nil {
			return p
		}
	}

	return nil
}
