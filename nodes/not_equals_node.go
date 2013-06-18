package nodes

type NotEqualsNode struct {
  *Node
}

func NotEquals(a, b interface{}) *ComparisonNode {
  return Comparison(&NotEqualsNode{&Node{a, b}})
}
