package nodes

// And node is a Binary node struct
type AndNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the And and other.
func (and *AndNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(and, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the And and other.
func (and *AndNode) And(other interface{}) *GroupingNode {
  return Grouping(And(and, other))
}

// Returns an Not node with and expression containing the
// And node.
func (and *AndNode) Not() *NotNode {
  return Not(and)
}

// AndNode factory method.
func And(left, right interface{}) *AndNode {
  and := new(AndNode)
  and.Left = left
  and.Right = right
  return and
}
