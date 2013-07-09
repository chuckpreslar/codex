package nodes

// Grouping node is a Unary node struct
type GroupingNode UnaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Grouping and other.
func (grouping *GroupingNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(grouping, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Grouping and other.
func (grouping *GroupingNode) And(other interface{}) *GroupingNode {
  return Grouping(And(grouping, other))
}

// Returns an Not node with and expression containing the
// Grouping node.
func (grouping *GroupingNode) Not() *NotNode {
  return Not(grouping)
}

// GroupingNode factory method.
func Grouping(expr interface{}) *GroupingNode {
  grouping := new(GroupingNode)
  grouping.Expr = expr
  return grouping
}
