package nodes

type DoesNotMatchNode struct {
  *Node
}

func DoesNotMatch(a, b interface{}) *DoesNotMatchNode {
  return &DoesNotMatchNode{&Node{a, b}}
}
