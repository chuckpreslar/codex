package nodes

type OrNode struct {
  *Node
}

func Or(comparison, other ComparisonInterface) *OrNode {
  return &OrNode{&Node{comparison, other}}
}
