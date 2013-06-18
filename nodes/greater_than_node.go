package nodes

type GreaterThanNode struct {
  *Node
}

func GreaterThan(a, b interface{}) *GreaterThanNode {
  return &GreaterThanNode{&Node{a, b}}
}
