package nodes

// AscendingNode is a UnaryNode struct
type AscendingNode UnaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Ascending and other.
func (asc *AscendingNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(asc, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Ascending and other.
func (asc *AscendingNode) And(other interface{}) *GroupingNode {
  return Grouping(And(asc, other))
}

// Returns an Ascending node with and expression containing the
// Ascending node.
func (asc *AscendingNode) Not() *NotNode {
  return Not(asc)
}

// AscendingNode factory method.
func Ascending(expr interface{}) *AscendingNode {
  asc := new(AscendingNode)
  asc.Expr = expr
  return asc
}
