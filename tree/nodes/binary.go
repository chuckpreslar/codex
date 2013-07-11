package nodes

// Binary node struct
type BinaryNode struct {
  Left  interface{} // Binary nodes left leaf.
  Right interface{} // Binary nodes right leaf.
}

type AsNode BinaryNode         // AsNode is a BinaryNode struct
type BetweenNode BinaryNode    // BetweenNode is a BinaryNode struct
type InnerJoinNode BinaryNode  // InnerJoinNode is a BinaryNode struct
type OuterJoinNode BinaryNode  // OuterJoinNode is a BinaryNode struct
type AssignmentNode BinaryNode // AssignmentNode is a BinaryNode struct

// AsNode factory method.
func As(left, right interface{}) (as *AsNode) {
  as = new(AsNode)
  as.Left = left
  as.Right = right
  return
}

// BetweenNode factory method.
func Between(left, right interface{}) (between *BetweenNode) {
  between = new(BetweenNode)
  between.Left = left
  between.Right = right
  return
}

// InnerJoinNode factory method.
func InnerJoin(left, right interface{}) (join *InnerJoinNode) {
  join = new(InnerJoinNode)
  join.Left = left
  join.Right = right
  return
}

// OuterJoinNode factory method.
func OuterJoin(left, right interface{}) (join *OuterJoinNode) {
  join = new(OuterJoinNode)
  join.Left = left
  join.Right = right
  return
}

// AssignmentNode factory method.
func Assignment(left, right interface{}) (assignment *AssignmentNode) {
  assignment = new(AssignmentNode)
  assignment.Left = left
  assignment.Right = right
  return
}
