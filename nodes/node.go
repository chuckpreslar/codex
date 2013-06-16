package nodes

type Node struct {
  left  interface{}
  right interface{}
}

func (node *Node) Left() interface{} {
  return node.left
}

func (node *Node) Right() interface{} {
  return node.right
}
