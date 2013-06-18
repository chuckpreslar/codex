package nodes

type LessThanNode struct {
  *Node
}

func LessThan(a, b interface{}) *ComparisonNode {
  return Comparison(&LessThanNode{&Node{a, b}})
}
