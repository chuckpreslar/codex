package nodes

// SumNode is a FunctionNode struct
type SumNode FunctionNode

// Returns and Equal node containing a reference to the
// function and other
func (sum *SumNode) Eq(other interface{}) *EqualNode {
  return Equal(sum, other)
}

// Returns and NotEqual node containing a reference to the
// function and other
func (sum *SumNode) Neq(other interface{}) *NotEqualNode {
  return NotEqual(sum, other)
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (sum *SumNode) Gt(other interface{}) *GreaterThanNode {
  return GreaterThan(sum, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (sum *SumNode) Gte(other interface{}) *GreaterThanOrEqualNode {
  return GreaterThanOrEqual(sum, other)
}

// Returns and LessThan node containing a reference to the
// function and other
func (sum *SumNode) Lt(other interface{}) *LessThanNode {
  return LessThan(sum, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (sum *SumNode) Lte(other interface{}) *LessThanOrEqualNode {
  return LessThanOrEqual(sum, other)
}

// Returns and Like node containing a reference to the
// function and other
func (sum *SumNode) Like(other interface{}) *LikeNode {
  return Like(sum, other)
}

// Returns and Unlike node containing a reference to the
// function and other
func (sum *SumNode) Unlike(other interface{}) *UnlikeNode {
  return Unlike(sum, other)
}

// Returns and Or node containing a reference to the
// function and other
func (sum *SumNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(sum, other))
}

// Returns and And node containing a reference to the
// function and other
func (sum *SumNode) And(other interface{}) *GroupingNode {
  return Grouping(And(sum, other))
}

// SumNode factory method.
func Sum(expressions ...interface{}) *SumNode {
  sum := new(SumNode)
  sum.Expressions = expressions
  return sum
}
