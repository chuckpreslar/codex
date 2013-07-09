package nodes

// GreaterThan node is a Binary node struct
type GreaterThanNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the GreaterThan and other.
func (gt *GreaterThanNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(gt, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the GreaterThan and other.
func (gt *GreaterThanNode) And(other interface{}) *GroupingNode {
  return Grouping(And(gt, other))
}

// Returns an Not node with and expression containing the
// GreaterThan node.
func (gt *GreaterThanNode) Not() *NotNode {
  return Not(gt)
}

func GreaterThan(left, right interface{}) *GreaterThanNode {
  gt := new(GreaterThanNode)
  gt.Left = left
  gt.Right = right
  return gt
}
