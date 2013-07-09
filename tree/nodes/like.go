package nodes

// Like node is a Binary node struct
type LikeNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Like and other.
func (like *LikeNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(like, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Like and other.
func (like *LikeNode) And(other interface{}) *GroupingNode {
  return Grouping(And(like, other))
}

// Returns an Not node with and expression containing the
// Like node.
func (like *LikeNode) Not() *NotNode {
  return Not(like)
}

func Like(left, right interface{}) *LikeNode {
  like := new(LikeNode)
  like.Left = left
  like.Right = right
  return like
}
