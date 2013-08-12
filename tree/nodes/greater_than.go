// Package nodes provides nodes to use in codex AST's.
package nodes

// GreaterThanNode is a BinaryNode struct
type GreaterThanNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the GreaterThan and other.
func (self *GreaterThanNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the GreaterThan and other.
func (self *GreaterThanNode) And(other interface{}) *GroupingNode {
  return Grouping(And(self, other))
}

// Returns an Not node with and expression containing the
// GreaterThan node.
func (self *GreaterThanNode) Not() *NotNode {
  return Not(self)
}

// GreaterThanNode factory method.
func GreaterThan(left, right interface{}) (gt *GreaterThanNode) {
  gt = new(GreaterThanNode)
  gt.Left = left
  gt.Right = right
  return
}
