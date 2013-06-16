package nodes

type LteNode struct {
  *Node
}

func (lte *LteNode) Or(other ComparisonInterface) OrInterface {
  return Or(lte, other)
}

func (lte *LteNode) Attribute() AttributeInterface {
  return lte.Left().(AttributeInterface)
}

func (lte *LteNode) Value() interface{} {
  return lte.Right()
}

func Lte(attribute AttributeInterface, value interface{}) *LteNode {
  return &LteNode{&Node{attribute, value}}
}
