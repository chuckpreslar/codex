package nodes

type EqualsNode struct {
  *Node
}

func Equals(a, b interface{}) *EqualsNode {
  return &EqualsNode{&Node{a, b}}
}
