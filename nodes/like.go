// Package nodes provides nodes to use in codex AST's.
package nodes

// LikeNode is a BinaryNode struct
type LikeNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Like and other.
func (self *LikeNode) Or(other interface{}) *GroupingNode {
	return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Like and other.
func (self *LikeNode) And(other interface{}) *GroupingNode {
	return Grouping(And(self, other))
}

// Returns an Not node with and expression containing the
// Like node.
func (self *LikeNode) Not() *NotNode {
	return Not(self)
}

// LikeNode factory method.
func Like(left, right interface{}) (like *LikeNode) {
	like = new(LikeNode)
	like.Left = left
	like.Right = right
	return
}
