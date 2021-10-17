package core

type Condition struct {
	Left Node
	Operator string
	Right Node
}

func (c Condition) String() string {
	// Unary condition
	if len(c.Operator) == 0 && c.Right == nil {
		return c.Left.String()
	}

	// Assignment condition
	if c.Left != nil && len(c.Operator) == 0 && c.Right != nil {
		return c.Left.String() + "; " + c.Right.String()
	}

	// Binary condition
	return c.Left.String() + " " + c.Operator + " " + c.Right.String()
}

