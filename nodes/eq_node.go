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

func Eq(value ComparableInterface, other interface{}) ComparisonInterface {
  return &EqNode{&Node{value, other}}
}
