package nodes

type GreaterThanNode struct {
  *Node
}

func GreaterThan(a, b interface{}) *ComparisonNode {
  return Comparison(&GreaterThanNode{&Node{a, b}})
}
