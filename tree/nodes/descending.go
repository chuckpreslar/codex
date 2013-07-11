package nodes

// DescendingNode is a UnaryNode struct
type DescendingNode UnaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Descending and other.
func (self *DescendingNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Descending and other.
func (self *DescendingNode) And(other interface{}) *GroupingNode {
  return Grouping(And(self, other))
}

// Returns an Descending node with and expression containing the
// Descending node.
func (self *DescendingNode) Not() *NotNode {
  return Not(self)
}

// DescendingNode factory method.
func Descending(expr interface{}) (descending *DescendingNode) {
  descending = new(DescendingNode)
  descending.Expr = expr
  return
}
