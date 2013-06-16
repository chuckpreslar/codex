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

func Gt(value ComparableInterface, other interface{}) ComparisonInterface {
  return &GtNode{&Node{value, other}}
}
