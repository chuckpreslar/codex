package nodes

type LessThanOrEqualsNode struct {
  *Node
}

func LessThanOrEquals(a, b interface{}) *ComparisonNode {
  return Comparison(&LessThanOrEqualsNode{&Node{a, b}})
}
