package nodes

type LimitNode struct {
  *Node
}

func Limit(take interface{}) *LimitNode {
  return &LimitNode{&Node{take, nil}}
}
