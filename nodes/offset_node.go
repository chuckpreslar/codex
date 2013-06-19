package nodes

type OffsetNode struct {
  *Node
}

func Offset(skip interface{}) *OffsetNode {
  return &OffsetNode{&Node{skip, nil}}
}
