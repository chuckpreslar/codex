package nodes

type AsNode struct {
  *Node
}

func As(a, b interface{}) *AsNode {
  return &AsNode{&Node{a, b}}
}
