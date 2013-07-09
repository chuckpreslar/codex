package nodes

// UnlikeNode is a BinaryNode struct
type UnlikeNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Unlike and other.
func (unlike *UnlikeNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(unlike, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Unlike and other.
func (unlike *UnlikeNode) And(other interface{}) *GroupingNode {
  return Grouping(And(unlike, other))
}

// Returns an Not node with and expression containing the
// Unlike node.
func (unlike *UnlikeNode) Not() *NotNode {
  return Not(unlike)
}

// UnlikeNode factory method.
func Unlike(left, right interface{}) *UnlikeNode {
  unlike := new(UnlikeNode)
  unlike.Left = left
  unlike.Right = right
  return unlike
}
