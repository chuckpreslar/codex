package nodes

type LtNode struct {
  *Node
}

func (lt *LtNode) Or(other ComparisonInterface) ComparisonInterface {
  return Or(lt, other)
}

func (lt *LtNode) Attribute() AttributeInterface {
  return lt.Left().(AttributeInterface)
}

func (lt *LtNode) Value() interface{} {
  return lt.Right()
}

func Lt(value ComparableInterface, other interface{}) ComparisonInterface {
  return &LtNode{&Node{value, other}}
}
