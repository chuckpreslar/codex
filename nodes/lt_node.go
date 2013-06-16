package nodes

type LtNode struct {
  *Node
}

func (lt *LtNode) Or(other ComparisonInterface) OrInterface {
  return Or(lt, other)
}

func (lt *LtNode) Attribute() AttributeInterface {
  return lt.Left().(AttributeInterface)
}

func (lt *LtNode) Value() interface{} {
  return lt.Right()
}

func (lt *LtNode) Left() AttributeInterface {
  return lt.Attribute()
}

func Lt(attribute AttributeInterface, value interface{}) *LtNode {
  return &LtNode{&Node{attribute, value}}
}
