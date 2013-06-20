package nodes

type FromNode struct {
  *Node
}

func From(a interface{}) *FromNode {
  return &FromNode{&Node{a, nil}}
}
