package nodes

type LessThanNode struct {
  *Node
}

func LessThan(a, b interface{}) *LessThanNode {
  return &LessThanNode{&Node{a, b}}
}
