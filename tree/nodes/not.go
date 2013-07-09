package nodes

// NotNode is a UnaryNode struct
type NotNode UnaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Not and other.
func (not *NotNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(not, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Not and other.
func (not *NotNode) And(other interface{}) *GroupingNode {
  return Grouping(And(not, other))
}

// Returns an Not node with and expression containing the
// Not node.
func (not *NotNode) Not() *NotNode {
  return Not(not)
}

// NotNode factory method.
func Not(expr interface{}) *NotNode {
  not := new(NotNode)
  not.Expr = expr
  return not
}
