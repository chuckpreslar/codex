package nodes

type CountNode FunctionNode

// Returns and Equal node containing a reference to the
// function and other
func (count *CountNode) Eq(other interface{}) *EqualNode {
  return Equal(count, other)
}

// Returns and NotEqual node containing a reference to the
// function and other
func (count *CountNode) Neq(other interface{}) *NotEqualNode {
  return NotEqual(count, other)
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (count *CountNode) Gt(other interface{}) *GreaterThanNode {
  return GreaterThan(count, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (count *CountNode) Gte(other interface{}) *GreaterThanOrEqualNode {
  return GreaterThanOrEqual(count, other)
}

// Returns and LessThan node containing a reference to the
// function and other
func (count *CountNode) Lt(other interface{}) *LessThanNode {
  return LessThan(count, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (count *CountNode) Lte(other interface{}) *LessThanOrEqualNode {
  return LessThanOrEqual(count, other)
}

// Returns and Like node containing a reference to the
// function and other
func (count *CountNode) Like(other interface{}) *LikeNode {
  return Like(count, other)
}

// Returns and Unlike node containing a reference to the
// function and other
func (count *CountNode) Unlike(other interface{}) *UnlikeNode {
  return Unlike(count, other)
}

// Returns and Or node containing a reference to the
// function and other
func (count *CountNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(count, other))
}

// Returns and And node containing a reference to the
// function and other
func (count *CountNode) And(other interface{}) *GroupingNode {
  return Grouping(And(count, other))
}

func Count(expressions ...interface{}) *CountNode {
  count := new(CountNode)
  count.Expressions = expressions
  return count
}
