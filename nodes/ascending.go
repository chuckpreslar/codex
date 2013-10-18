// Package nodes provides nodes to use in codex AST's.
package nodes

// AscendingNode is a UnaryNode struct.
type AscendingNode UnaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Ascending and other.
func (self *AscendingNode) Or(other interface{}) *GroupingNode {
	return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Ascending and other.
func (self *AscendingNode) And(other interface{}) *GroupingNode {
	return Grouping(And(self, other))
}

// Returns an Ascending node with and expression containing the
// Ascending node.
func (self *AscendingNode) Not() *NotNode {
	return Not(self)
}

// AscendingNode factory method.
func Ascending(expr interface{}) (ascending *AscendingNode) {
	ascending = new(AscendingNode)
	ascending.Expr = expr
	return
}
