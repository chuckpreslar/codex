package nodes

// Binary node struct
type BinaryNode struct {
  Left  interface{} // Binary nodes left leaf.
  Right interface{} // Binary nodes right leaf.
}

type AsNode BinaryNode         // As node is a Binary node struct
type BetweenNode BinaryNode    // Between node is a Binary node struct
type InnerJoinNode BinaryNode  // InnerJoin node is a Binary node struct
type OuterJoinNode BinaryNode  // OuterJoin node is a Binary node struct
type AssignmentNode BinaryNode // Assignment node is a Binary node struct

func As(left, right interface{}) *AsNode {
  as := new(AsNode)
  as.Left = left
  as.Right = right
  return as
}

func Between(left, right interface{}) *BetweenNode {
  between := new(BetweenNode)
  between.Left = left
  between.Right = right
  return between
}

func InnerJoin(left, right interface{}) *InnerJoinNode {
  join := new(InnerJoinNode)
  join.Left = left
  join.Right = right
  return join
}

func OuterJoin(left, right interface{}) *OuterJoinNode {
  join := new(OuterJoinNode)
  join.Left = left
  join.Right = right
  return join
}

func Assignment(left, right interface{}) *AssignmentNode {
  assignment := new(AssignmentNode)
  assignment.Left = left
  assignment.Right = right
  return assignment
}
