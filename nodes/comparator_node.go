package nodes

type ComparatorNode struct {
  BaseNode
}

func (comparator ComparatorNode) Or(other ComparatorNode) ComparatorNode {
  comparator.right = other
  return comparator
}
