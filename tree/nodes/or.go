package nodes

// OrNode is a BinaryNode struct
type OrNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Or and other.
func (or *OrNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(or, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Or and other.
func (or *OrNode) And(other interface{}) *GroupingNode {
  return Grouping(And(or, other))
}

// Returns an Not node with and expression containing the
// Or node.
func (or *OrNode) Not() *NotNode {
  return Not(or)
}

// OrNode factory method.
func Or(left, right interface{}) *OrNode {
  or := new(OrNode)
  or.Left = left
  or.Right = right
  return or
}
