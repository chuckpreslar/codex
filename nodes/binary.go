// Package nodes provides nodes to use in codex AST's.
package nodes

// Binary node struct.
type BinaryNode struct {
  Left  interface{} // Binary nodes left leaf.
  Right interface{} // Binary nodes right leaf.
}

type AsNode BinaryNode               // AsNode is a BinaryNode struct.
type BetweenNode BinaryNode          // BetweenNode is a BinaryNode struct.
type InnerJoinNode BinaryNode        // InnerJoinNode is a BinaryNode struct.
type OuterJoinNode BinaryNode        // OuterJoinNode is a BinaryNode struct.
type AssignmentNode BinaryNode       // AssignmentNode is a BinaryNode struct.
type UnionNode BinaryNode            // UnionNode is a BinaryNode struct.
type IntersectNode BinaryNode        // IntersectNode is a BinaryNode struct.
type ExceptNode BinaryNode           // ExceptNode is a BinaryNode struct.
type UnexistingColumnNode BinaryNode // UnexistingColumnNode is a BinaryNode struct.

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

// UnionNode factory method.
func Union(left, right interface{}) (union *UnionNode) {
  union = new(UnionNode)
  union.Left = left
  union.Right = right
  return
}

// IntersectNode factory method.
func Intersect(left, right interface{}) (intersect *IntersectNode) {
  intersect = new(IntersectNode)
  intersect.Left = left
  intersect.Right = right
  return
}

// ExceptNode factory method.
func Except(left, right interface{}) (except *ExceptNode) {
  except = new(ExceptNode)
  except.Left = left
  except.Right = right
  return
}

// UnexistingColumnNode factory method.
func UnexistingColumn(left, right interface{}) (column *UnexistingColumnNode) {
  column = new(UnexistingColumnNode)
  column.Left = left
  column.Right = right
  return
}
