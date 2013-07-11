package nodes

// AndNode is a BinaryNode struct
type AndNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the And and other.
func (self *AndNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the And and other.
func (self *AndNode) And(other interface{}) *GroupingNode {
  return Grouping(And(self, other))
}

// Returns an Not node with and expression containing the
// And node.
func (self *AndNode) Not() *NotNode {
  return Not(self)
}

// AndNode factory method.
func And(left, right interface{}) (and *AndNode) {
  and = new(AndNode)
  and.Left = left
  and.Right = right
  return
}
