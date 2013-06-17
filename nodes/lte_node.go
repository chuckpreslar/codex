package nodes

type LteNode struct {
  *Node
}

func (lte *LteNode) Or(other ComparisonInterface) ComparisonInterface {
  return Or(lte, other)
}

func (lte *LteNode) Attribute() AttributeInterface {
  return lte.Left.(AttributeInterface)
}

func (lte *LteNode) Value() interface{} {
  return lte.Right
}

func Lte(value ComparableInterface, other interface{}) *LteNode {
  return &LteNode{&Node{value, other}}
}
