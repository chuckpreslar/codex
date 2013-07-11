package nodes

type DescendingNode UnaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Descending and other.
func (desc *DescendingNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(desc, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Descending and other.
func (desc *DescendingNode) And(other interface{}) *GroupingNode {
  return Grouping(And(desc, other))
}

// Returns an Descending node with and expression containing the
// Descending node.
func (desc *DescendingNode) Not() *NotNode {
  return Not(desc)
}

// DescendingNode factory method.
func Descending(expr interface{}) *DescendingNode {
  desc := new(DescendingNode)
  desc.Expr = expr
  return desc
}