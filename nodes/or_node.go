package nodes

type OrNode struct {
  *Node
}

func (or *OrNode) Left() ComparisonInterface {
  return or.Node.Left().(ComparisonInterface)
}

func (or *OrNode) Right() ComparisonInterface {
  return or.Node.Right().(ComparisonInterface)
}

func Or(comparison, other ComparisonInterface) *OrNode {
  return &OrNode{&Node{comparison, other}}
}
