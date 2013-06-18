package nodes

type EqualsNode struct {
  *Node
}

func Equals(a, b interface{}) *ComparisonNode {
  return Comparison(&EqualsNode{&Node{a, b}})
}
