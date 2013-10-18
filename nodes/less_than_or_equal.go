// Package nodes provides nodes to use in codex AST's.
package nodes

// LessThanOrEqualNode is a BinaryNode struct
type LessThanOrEqualNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the LessThanOrEqual and other.
func (self *LessThanOrEqualNode) Or(other interface{}) *GroupingNode {
	return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the LessThanOrEqual and other.
func (self *LessThanOrEqualNode) And(other interface{}) *GroupingNode {
	return Grouping(And(self, other))
}

// Returns an Not node with and expression containing the
// LessThanOrEqual node.
func (self *LessThanOrEqualNode) Not() *NotNode {
	return Not(self)
}

// LessThanOrEqualNode factory method.
func LessThanOrEqual(left, right interface{}) (lte *LessThanOrEqualNode) {
	lte = new(LessThanOrEqualNode)
	lte.Left = left
	lte.Right = right
	return
}
