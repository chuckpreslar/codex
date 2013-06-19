package nodes

type OnNode struct {
  *Node
}

func On(a, b interface{}) *OnNode {
  return &OnNode{&Node{a, b}}
}
