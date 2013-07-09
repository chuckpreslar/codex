package nodes

// MinimumNode is a FunctionNode struct
type MinimumNode FunctionNode

// Returns and Equal node containing a reference to the
// function and other
func (min *MinimumNode) Eq(other interface{}) *EqualNode {
  return Equal(min, other)
}

// Returns and NotEqual node containing a reference to the
// function and other
func (min *MinimumNode) Neq(other interface{}) *NotEqualNode {
  return NotEqual(min, other)
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (min *MinimumNode) Gt(other interface{}) *GreaterThanNode {
  return GreaterThan(min, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (min *MinimumNode) Gte(other interface{}) *GreaterThanOrEqualNode {
  return GreaterThanOrEqual(min, other)
}

// Returns and LessThan node containing a reference to the
// function and other
func (min *MinimumNode) Lt(other interface{}) *LessThanNode {
  return LessThan(min, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (min *MinimumNode) Lte(other interface{}) *LessThanOrEqualNode {
  return LessThanOrEqual(min, other)
}

// Returns and Like node containing a reference to the
// function and other
func (min *MinimumNode) Like(other interface{}) *LikeNode {
  return Like(min, other)
}

// Returns and Unlike node containing a reference to the
// function and other
func (min *MinimumNode) Unlike(other interface{}) *UnlikeNode {
  return Unlike(min, other)
}

// Returns and Or node containing a reference to the
// function and other
func (min *MinimumNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(min, other))
}

// Returns and And node containing a reference to the
// function and other
func (min *MinimumNode) And(other interface{}) *GroupingNode {
  return Grouping(And(min, other))
}

// MinimumNode factory method.
func Minimum(expressions ...interface{}) *MinimumNode {
  min := new(MinimumNode)
  min.Expressions = expressions
  return min
}
