// Package nodes provides nodes to use in codex AST's.
package nodes

// UnlikeNode is a BinaryNode struct
type UnlikeNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Unlike and other.
func (self *UnlikeNode) Or(other interface{}) *GroupingNode {
	return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Unlike and other.
func (self *UnlikeNode) And(other interface{}) *GroupingNode {
	return Grouping(And(self, other))
}

// Returns an Not node with and expression containing the
// Unlike node.
func (self *UnlikeNode) Not() *NotNode {
	return Not(self)
}

// UnlikeNode factory method.
func Unlike(left, right interface{}) (unlike *UnlikeNode) {
	unlike = new(UnlikeNode)
	unlike.Left = left
	unlike.Right = right
	return
}
