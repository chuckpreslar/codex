// Package nodes provides nodes to use in codex AST's.
package nodes

// MinimumNode is a FunctionNode struct
type MinimumNode FunctionNode

// Returns and Equal node containing a reference to the
// function and other
func (self *MinimumNode) Eq(other interface{}) *EqualNode {
	return Equal(self, other)
}

// Returns and NotEqual node containing a reference to the
// function and other
func (self *MinimumNode) Neq(other interface{}) *NotEqualNode {
	return NotEqual(self, other)
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (self *MinimumNode) Gt(other interface{}) *GreaterThanNode {
	return GreaterThan(self, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (self *MinimumNode) Gte(other interface{}) *GreaterThanOrEqualNode {
	return GreaterThanOrEqual(self, other)
}

// Returns and LessThan node containing a reference to the
// function and other
func (self *MinimumNode) Lt(other interface{}) *LessThanNode {
	return LessThan(self, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (self *MinimumNode) Lte(other interface{}) *LessThanOrEqualNode {
	return LessThanOrEqual(self, other)
}

// Returns and Like node containing a reference to the
// function and other
func (self *MinimumNode) Like(other interface{}) *LikeNode {
	return Like(self, other)
}

// Returns and Unlike node containing a reference to the
// function and other
func (self *MinimumNode) Unlike(other interface{}) *UnlikeNode {
	return Unlike(self, other)
}

// Returns and Or node containing a reference to the
// function and other
func (self *MinimumNode) Or(other interface{}) *GroupingNode {
	return Grouping(Or(self, other))
}

// Returns and And node containing a reference to the
// function and other
func (self *MinimumNode) And(other interface{}) *GroupingNode {
	return Grouping(And(self, other))
}

// MinimumNode factory method.
func Minimum(expressions ...interface{}) (minimum *MinimumNode) {
	minimum = new(MinimumNode)
	minimum.Expressions = expressions
	return
}
