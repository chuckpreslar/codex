// Package nodes provides nodes to use in codex AST's.
package nodes

// GreaterThanOrEqualNode is a BinaryNode struct
type GreaterThanOrEqualNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the GreaterThanOrEqual and other.
func (self *GreaterThanOrEqualNode) Or(other interface{}) *GroupingNode {
	return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the GreaterThanOrEqual and other.
func (self *GreaterThanOrEqualNode) And(other interface{}) *GroupingNode {
	return Grouping(And(self, other))
}

// Returns an Not node with and expression containing the
// GreaterThanOrEqual node.
func (self *GreaterThanOrEqualNode) Not() *NotNode {
	return Not(self)
}

// GreaterThanOrEqualNode factory method.
func GreaterThanOrEqual(left, right interface{}) (gte *GreaterThanOrEqualNode) {
	gte = new(GreaterThanOrEqualNode)
	gte.Left = left
	gte.Right = right
	return
}
