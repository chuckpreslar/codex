package nodes

type AverageNode FunctionNode

// Returns and Equal node containing a reference to the
// function and other
func (avg *AverageNode) Eq(other interface{}) *EqualNode {
  return Equal(avg, other)
}

// Returns and NotEqual node containing a reference to the
// function and other
func (avg *AverageNode) Neq(other interface{}) *NotEqualNode {
  return NotEqual(avg, other)
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (avg *AverageNode) Gt(other interface{}) *GreaterThanNode {
  return GreaterThan(avg, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (avg *AverageNode) Gte(other interface{}) *GreaterThanOrEqualNode {
  return GreaterThanOrEqual(avg, other)
}

// Returns and LessThan node containing a reference to the
// function and other
func (avg *AverageNode) Lt(other interface{}) *LessThanNode {
  return LessThan(avg, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (avg *AverageNode) Lte(other interface{}) *LessThanOrEqualNode {
  return LessThanOrEqual(avg, other)
}

// Returns and Like node containing a reference to the
// function and other
func (avg *AverageNode) Like(other interface{}) *LikeNode {
  return Like(avg, other)
}

// Returns and Unlike node containing a reference to the
// function and other
func (avg *AverageNode) Unlike(other interface{}) *UnlikeNode {
  return Unlike(avg, other)
}

// Returns and Or node containing a reference to the
// function and other
func (avg *AverageNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(avg, other))
}

// Returns and And node containing a reference to the
// function and other
func (avg *AverageNode) And(other interface{}) *GroupingNode {
  return Grouping(And(avg, other))
}

func Average(expressions ...interface{}) *AverageNode {
  avg := new(AverageNode)
  avg.Expressions = expressions
  return avg
}
