package nodes

// NotNode is a UnaryNode struct
type NotNode UnaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Not and other.
func (self *NotNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Not and other.
func (self *NotNode) And(other interface{}) *GroupingNode {
  return Grouping(And(self, other))
}

// Returns an Not node with and expression containing the
// Not node.
func (self *NotNode) Not() *NotNode {
  return Not(self)
}

// NotNode factory method.
func Not(expr interface{}) (not *NotNode) {
  not = new(NotNode)
  not.Expr = expr
  return
}
