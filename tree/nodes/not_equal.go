package nodes

// NotEqualNode is a BinaryNode struct
type NotEqualNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the NotEqual and other.
func (neq *NotEqualNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(neq, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the NotEqual and other.
func (neq *NotEqualNode) And(other interface{}) *GroupingNode {
  return Grouping(And(neq, other))
}

// Returns an Not node with and expression containing the
// NotEqual node.
func (neq *NotEqualNode) Not() *NotNode {
  return Not(neq)
}

// NotEqualNode factory method.
func NotEqual(left, right interface{}) *NotEqualNode {
  neq := new(NotEqualNode)
  neq.Left = left
  neq.Right = right
  return neq
}
