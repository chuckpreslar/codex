package nodes

type NeqNode struct {
  *Node
}

func (neq *NeqNode) Or(other ComparisonInterface) ComparisonInterface {
  return Or(neq, other)
}

func (neq *NeqNode) Attribute() AttributeInterface {
  return neq.Left.(AttributeInterface)
}

func (neq *NeqNode) Value() interface{} {
  return neq.Right
}

func Neq(value ComparableInterface, other interface{}) *NeqNode {
  return &NeqNode{&Node{value, other}}
}
