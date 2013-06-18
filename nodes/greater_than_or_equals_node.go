package nodes

type GreaterThanOrEqualsNode struct {
  *Node
}

func GreaterThanOrEquals(a, b interface{}) *ComparisonNode {
  return Comparison(&GreaterThanOrEqualsNode{&Node{a, b}})
}
