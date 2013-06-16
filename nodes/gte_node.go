package nodes

type GteNode struct {
  *Node
}

func (gte *GteNode) Or(other ComparisonInterface) OrInterface {
  return Or(gte, other)
}

func (gte *GteNode) Attribute() AttributeInterface {
  return gte.Left().(AttributeInterface)
}

func (gte *GteNode) Value() interface{} {
  return gte.Right()
}

func Gte(attribute AttributeInterface, value interface{}) *GteNode {
  return &GteNode{&Node{attribute, value}}
}
