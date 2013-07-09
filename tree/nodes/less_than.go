package nodes

// LessThan node is a Binary node struct
type LessThanNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the LessThan and other.
func (lt *LessThanNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(lt, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the LessThanNode and other.
func (lt *LessThanNode) And(other interface{}) *GroupingNode {
  return Grouping(And(lt, other))
}

// Returns an Not node with and expression containing the
// LessThan node.
func (lt *LessThanNode) Not() *NotNode {
  return Not(lt)
}

func LessThan(left, right interface{}) *LessThanNode {
  lt := new(LessThanNode)
  lt.Left = left
  lt.Right = right
  return lt
}
