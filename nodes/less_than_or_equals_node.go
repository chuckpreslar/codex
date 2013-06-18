package nodes

type LessThanOrEqualsNode struct {
  *Node
}

func LessThanOrEquals(a, b interface{}) *LessThanOrEqualsNode {
  return &LessThanOrEqualsNode{&Node{a, b}}
}
