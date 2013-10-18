// Package nodes provides nodes to use in codex AST's.
package nodes

// EqualNode is a BinaryNode struct
type EqualNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Equal and other.
func (self *EqualNode) Or(other interface{}) *GroupingNode {
	return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Equal and other.
func (self *EqualNode) And(other interface{}) *GroupingNode {
	return Grouping(And(self, other))
}

// Returns an Not node with and expression containing the
// Equal node.
func (self *EqualNode) Not() *NotNode {
	return Not(self)
}

// Equal factory method.
func Equal(left, right interface{}) (eq *EqualNode) {
	eq = new(EqualNode)
	eq.Left = left
	eq.Right = right
	return
}
