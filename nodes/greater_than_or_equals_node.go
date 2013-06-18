package nodes

type GreaterThanOrEqualsNode struct {
  *Node
}

func GreaterThanOrEquals(a, b interface{}) *GreaterThanOrEqualsNode {
  return &GreaterThanOrEqualsNode{&Node{a, b}}
}
