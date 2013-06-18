package nodes

type NotEqualsNode struct {
  *Node
}

func NotEquals(a, b interface{}) *NotEqualsNode {
  return &NotEqualsNode{&Node{a, b}}
}
