package nodes

type InnerJoinNode struct {
  *Node
}

func InnerJoin(a interface{}) *InnerJoinNode {
  return &InnerJoinNode{&Node{a, nil}}
}
