package nodes

type NeqNode struct {
  *Node
}

func (neq *NeqNode) Or(other ComparisonInterface) OrInterface {
  return Or(neq, other)
}

func (neq *NeqNode) Attribute() AttributeInterface {
  return neq.Left().(AttributeInterface)
}

func (neq *NeqNode) Value() interface{} {
  return neq.Right()
}

func Neq(attribute AttributeInterface, value interface{}) *NeqNode {
  return &NeqNode{&Node{attribute, value}}
}
