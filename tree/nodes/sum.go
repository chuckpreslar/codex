package nodes

// SumNode is a FunctionNode struct
type SumNode FunctionNode

// Returns and Equal node containing a reference to the
// function and other
func (self *SumNode) Eq(other interface{}) *EqualNode {
  return Equal(self, other)
}

// Returns and NotEqual node containing a reference to the
// function and other
func (self *SumNode) Neq(other interface{}) *NotEqualNode {
  return NotEqual(self, other)
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (self *SumNode) Gt(other interface{}) *GreaterThanNode {
  return GreaterThan(self, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (self *SumNode) Gte(other interface{}) *GreaterThanOrEqualNode {
  return GreaterThanOrEqual(self, other)
}

// Returns and LessThan node containing a reference to the
// function and other
func (self *SumNode) Lt(other interface{}) *LessThanNode {
  return LessThan(self, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (self *SumNode) Lte(other interface{}) *LessThanOrEqualNode {
  return LessThanOrEqual(self, other)
}

// Returns and Like node containing a reference to the
// function and other
func (self *SumNode) Like(other interface{}) *LikeNode {
  return Like(self, other)
}

// Returns and Unlike node containing a reference to the
// function and other
func (self *SumNode) Unlike(other interface{}) *UnlikeNode {
  return Unlike(self, other)
}

// Returns and Or node containing a reference to the
// function and other
func (self *SumNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(self, other))
}

// Returns and And node containing a reference to the
// function and other
func (self *SumNode) And(other interface{}) *GroupingNode {
  return Grouping(And(self, other))
}

// SumNode factory method.
func Sum(expressions ...interface{}) (sum *SumNode) {
  sum = new(SumNode)
  sum.Expressions = expressions
  return
}
