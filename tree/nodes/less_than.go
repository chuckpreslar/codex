// Package nodes provides nodes to use in codex AST's.
package nodes

// LessThanNode is a BinaryNode struct
type LessThanNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the LessThan and other.
func (self *LessThanNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the LessThanNode and other.
func (self *LessThanNode) And(other interface{}) *GroupingNode {
  return Grouping(And(self, other))
}

// Returns an Not node with and expression containing the
// LessThan node.
func (self *LessThanNode) Not() *NotNode {
  return Not(self)
}

// LessThanNode factory method.
func LessThan(left, right interface{}) (lt *LessThanNode) {
  lt = new(LessThanNode)
  lt.Left = left
  lt.Right = right
  return
}
