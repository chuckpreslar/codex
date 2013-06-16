package nodes

type GtNode struct {
  *Node
}

func (gt *GtNode) Or(other ComparisonInterface) OrInterface {
  return Or(gt, other)
}

func (gt *GtNode) Attribute() AttributeInterface {
  return gt.Left().(AttributeInterface)
}

func (gt *GtNode) Value() interface{} {
  return gt.Right()
}

func Gt(attribute AttributeInterface, value interface{}) *GtNode {
  return &GtNode{&Node{attribute, value}}
}
