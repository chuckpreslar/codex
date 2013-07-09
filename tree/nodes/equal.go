package nodes

// Equal node is a Binary node struct
type EqualNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Equal and other.
func (eq *EqualNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(eq, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Equal and other.
func (eq *EqualNode) And(other interface{}) *GroupingNode {
  return Grouping(And(eq, other))
}

// Returns an Not node with and expression containing the
// Equal node.
func (eq *EqualNode) Not() *NotNode {
  return Not(eq)
}

func Equal(left, right interface{}) *EqualNode {
  eq := new(EqualNode)
  eq.Left = left
  eq.Right = right
  return eq
}
