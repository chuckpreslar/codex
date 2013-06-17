package nodes

type EqNode struct {
  *Node
}

func (eq *EqNode) Or(other ComparisonInterface) ComparisonInterface {
  return Or(eq, other)
}

func (eq *EqNode) Attribute() AttributeInterface {
  return eq.Left.(AttributeInterface)
}

func (eq *EqNode) Value() interface{} {
  return eq.Right
}

func Eq(value ComparableInterface, other interface{}) *EqNode {
  return &EqNode{&Node{value, other}}
}
