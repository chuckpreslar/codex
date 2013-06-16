package nodes

type EqNode struct {
  *Node
}

func (eq *EqNode) Or(other ComparisonInterface) OrInterface {
  return Or(eq, other)
}

func (eq *EqNode) Attribute() AttributeInterface {
  return eq.Left().(AttributeInterface)
}

func (eq *EqNode) Value() interface{} {
  return eq.Right()
}

func Eq(attribute AttributeInterface, value interface{}) *EqNode {
  return &EqNode{&Node{attribute, value}}
}
