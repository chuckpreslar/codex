// Package nodes provides nodes to use in codex AST's.
package nodes

// Grouping node is a Unary node struct
type GroupingNode UnaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Grouping and other.
func (self *GroupingNode) Or(other interface{}) *GroupingNode {
	return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Grouping and other.
func (self *GroupingNode) And(other interface{}) *GroupingNode {
	return Grouping(And(self, other))
}

// Returns an Not node with and expression containing the
// Grouping node.
func (self *GroupingNode) Not() *NotNode {
	return Not(self)
}

// GroupingNode factory method.
func Grouping(expr interface{}) (grouping *GroupingNode) {
	grouping = new(GroupingNode)
	grouping.Expr = expr
	return
}
