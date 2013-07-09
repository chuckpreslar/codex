package nodes

// GreaterThanOrEqualNode is a BinaryNode struct
type GreaterThanOrEqualNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the GreaterThanOrEqual and other.
func (gte *GreaterThanOrEqualNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(gte, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the GreaterThanOrEqual and other.
func (gte *GreaterThanOrEqualNode) And(other interface{}) *GroupingNode {
  return Grouping(And(gte, other))
}

// Returns an Not node with and expression containing the
// GreaterThanOrEqual node.
func (gte *GreaterThanOrEqualNode) Not() *NotNode {
  return Not(gte)
}

// GreaterThanOrEqualNode factory method.
func GreaterThanOrEqual(left, right interface{}) *GreaterThanOrEqualNode {
  gte := new(GreaterThanOrEqualNode)
  gte.Left = left
  gte.Right = right
  return gte
}
