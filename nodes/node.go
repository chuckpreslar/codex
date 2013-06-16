package nodes

type Node struct {
  left  interface{}
  right interface{}
}

func (node *Node) GetLeft() interface{} {
  return node.left
}

func (node *Node) GetRight() interface{} {
  return node.right
}
