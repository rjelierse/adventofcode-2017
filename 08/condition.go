package registers

const (
	CompareEquals       Comparator = "=="
	CompareNotEquals               = "!="
	CompareGreater                 = ">"
	CompareGreaterEqual            = ">="
	CompareLower                   = "<"
	CompareLowerEqual              = "<="
)

type Comparator string

type Condition struct {
	Register   string
	Comparator Comparator
	Value      int
}

func (c *Condition) Valid() bool {
	value := GetValue(c.Register)

	switch c.Comparator {
	case CompareEquals:
		return value == c.Value
	case CompareNotEquals:
		return value != c.Value
	case CompareLower:
		return value < c.Value
	case CompareLowerEqual:
		return value <= c.Value
	case CompareGreaterEqual:
		return value >= c.Value
	case CompareGreater:
		return value > c.Value
	}

	return false
}
