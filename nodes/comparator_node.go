package nodes

type ComparatorNode BaseNode

func (comparator ComparatorNode) Or(other ComparatorNode) ComparatorNode {
  comparator.right = other
  return comparator
}
