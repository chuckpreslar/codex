package nodes

// CountNode is a FunctionNode struct
type CountNode FunctionNode

// Returns and Equal node containing a reference to the
// function and other
func (self *CountNode) Eq(other interface{}) *EqualNode {
  return Equal(self, other)
}

// Returns and NotEqual node containing a reference to the
// function and other
func (self *CountNode) Neq(other interface{}) *NotEqualNode {
  return NotEqual(self, other)
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (self *CountNode) Gt(other interface{}) *GreaterThanNode {
  return GreaterThan(self, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (self *CountNode) Gte(other interface{}) *GreaterThanOrEqualNode {
  return GreaterThanOrEqual(self, other)
}

// Returns and LessThan node containing a reference to the
// function and other
func (self *CountNode) Lt(other interface{}) *LessThanNode {
  return LessThan(self, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (self *CountNode) Lte(other interface{}) *LessThanOrEqualNode {
  return LessThanOrEqual(self, other)
}

// Returns and Like node containing a reference to the
// function and other
func (self *CountNode) Like(other interface{}) *LikeNode {
  return Like(self, other)
}

// Returns and Unlike node containing a reference to the
// function and other
func (self *CountNode) Unlike(other interface{}) *UnlikeNode {
  return Unlike(self, other)
}

// Returns and Or node containing a reference to the
// function and other
func (self *CountNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(self, other))
}

// Returns and And node containing a reference to the
// function and other
func (self *CountNode) And(other interface{}) *GroupingNode {
  return Grouping(And(self, other))
}

// CountNode factory method.
func Count(expressions ...interface{}) (count *CountNode) {
  count = new(CountNode)
  count.Expressions = expressions
  return
}
