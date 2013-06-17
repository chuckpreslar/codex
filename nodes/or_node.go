package nodes

type OrNode struct {
  *Node
}

func (or *OrNode) Or(other ComparisonInterface) ComparisonInterface {
  return Or(or, other)
}

func (or *OrNode) Attribute() AttributeInterface {
  return or.Left.(AttributeInterface)
}

func (or *OrNode) Value() interface{} {
  return or.Right
}

func Or(comparison, other ComparisonInterface) *OrNode {
  return &OrNode{&Node{comparison, other}}
}
