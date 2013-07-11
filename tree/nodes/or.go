package nodes

// OrNode is a BinaryNode struct
type OrNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Or and other.
func (self *OrNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Or and other.
func (self *OrNode) And(other interface{}) *GroupingNode {
  return Grouping(And(self, other))
}

// Returns an Not node with and expression containing the
// Or node.
func (self *OrNode) Not() *NotNode {
  return Not(self)
}

// OrNode factory method.
func Or(left, right interface{}) (or *OrNode) {
  or = new(OrNode)
  or.Left = left
  or.Right = right
  return
}
