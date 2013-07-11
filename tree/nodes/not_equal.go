package nodes

// NotEqualNode is a BinaryNode struct
type NotEqualNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the NotEqual and other.
func (self *NotEqualNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the NotEqual and other.
func (self *NotEqualNode) And(other interface{}) *GroupingNode {
  return Grouping(And(self, other))
}

// Returns an Not node with and expression containing the
// NotEqual node.
func (self *NotEqualNode) Not() *NotNode {
  return Not(self)
}

// NotEqualNode factory method.
func NotEqual(left, right interface{}) (neq *NotEqualNode) {
  neq = new(NotEqualNode)
  neq.Left = left
  neq.Right = right
  return
}
