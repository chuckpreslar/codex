package nodes

// MaximumNode is a FunctionNode struct
type MaximumNode FunctionNode

// Returns and Equal node containing a reference to the
// function and other
func (max *MaximumNode) Eq(other interface{}) *EqualNode {
  return Equal(max, other)
}

// Returns and NotEqual node containing a reference to the
// function and other
func (max *MaximumNode) Neq(other interface{}) *NotEqualNode {
  return NotEqual(max, other)
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (max *MaximumNode) Gt(other interface{}) *GreaterThanNode {
  return GreaterThan(max, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (max *MaximumNode) Gte(other interface{}) *GreaterThanOrEqualNode {
  return GreaterThanOrEqual(max, other)
}

// Returns and LessThan node containing a reference to the
// function and other
func (max *MaximumNode) Lt(other interface{}) *LessThanNode {
  return LessThan(max, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (max *MaximumNode) Lte(other interface{}) *LessThanOrEqualNode {
  return LessThanOrEqual(max, other)
}

// Returns and Like node containing a reference to the
// function and other
func (max *MaximumNode) Like(other interface{}) *LikeNode {
  return Like(max, other)
}

// Returns and Unlike node containing a reference to the
// function and other
func (max *MaximumNode) Unlike(other interface{}) *UnlikeNode {
  return Unlike(max, other)
}

// Returns and Or node containing a reference to the
// function and other
func (max *MaximumNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(max, other))
}

// Returns and And node containing a reference to the
// function and other
func (max *MaximumNode) And(other interface{}) *GroupingNode {
  return Grouping(And(max, other))
}

// MaximumNode factory method.
func Maximum(expressions ...interface{}) *MaximumNode {
  max := new(MaximumNode)
  max.Expressions = expressions
  return max
}
