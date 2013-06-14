package nodes

type ComparatorNode struct {
  BaseNode
}

func (comparator ComparatorNode) Or(other Node) (or OrNode) {
  or = OrNode{BaseNode{comparator, other}}
  return
}
