package nodes

// LessThanOrEqualNode is a BinaryNode struct
type LessThanOrEqualNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the LessThanOrEqual and other.
func (lte *LessThanOrEqualNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(lte, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the LessThanOrEqual and other.
func (lte *LessThanOrEqualNode) And(other interface{}) *GroupingNode {
  return Grouping(And(lte, other))
}

// Returns an Not node with and expression containing the
// LessThanOrEqual node.
func (lte *LessThanOrEqualNode) Not() *NotNode {
  return Not(lte)
}

// LessThanOrEqualNode factory method.
func LessThanOrEqual(left, right interface{}) *LessThanOrEqualNode {
  lte := new(LessThanOrEqualNode)
  lte.Left = left
  lte.Right = right
  return lte
}
