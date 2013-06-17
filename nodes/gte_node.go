package nodes

type GteNode struct {
  *Node
}

func (gte *GteNode) Or(other ComparisonInterface) ComparisonInterface {
  return Or(gte, other)
}

func (gte *GteNode) Attribute() AttributeInterface {
  return gte.Left.(AttributeInterface)
}

func (gte *GteNode) Value() interface{} {
  return gte.Right
}

func Gte(value ComparableInterface, other interface{}) *GteNode {
  return &GteNode{&Node{value, other}}
}
