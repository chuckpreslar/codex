package nodes

// MaximumNode is a FunctionNode struct
type MaximumNode FunctionNode

// Returns and Equal node containing a reference to the
// function and other
func (self *MaximumNode) Eq(other interface{}) *EqualNode {
  return Equal(self, other)
}

// Returns and NotEqual node containing a reference to the
// function and other
func (self *MaximumNode) Neq(other interface{}) *NotEqualNode {
  return NotEqual(self, other)
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (self *MaximumNode) Gt(other interface{}) *GreaterThanNode {
  return GreaterThan(self, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (self *MaximumNode) Gte(other interface{}) *GreaterThanOrEqualNode {
  return GreaterThanOrEqual(self, other)
}

// Returns and LessThan node containing a reference to the
// function and other
func (self *MaximumNode) Lt(other interface{}) *LessThanNode {
  return LessThan(self, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (self *MaximumNode) Lte(other interface{}) *LessThanOrEqualNode {
  return LessThanOrEqual(self, other)
}

// Returns and Like node containing a reference to the
// function and other
func (self *MaximumNode) Like(other interface{}) *LikeNode {
  return Like(self, other)
}

// Returns and Unlike node containing a reference to the
// function and other
func (self *MaximumNode) Unlike(other interface{}) *UnlikeNode {
  return Unlike(self, other)
}

// Returns and Or node containing a reference to the
// function and other
func (self *MaximumNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(self, other))
}

// Returns and And node containing a reference to the
// function and other
func (self *MaximumNode) And(other interface{}) *GroupingNode {
  return Grouping(And(self, other))
}

// MaximumNode factory method.
func Maximum(expressions ...interface{}) (maximum *MaximumNode) {
  maximum = new(MaximumNode)
  maximum.Expressions = expressions
  return
}
