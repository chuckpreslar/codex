package nodes

type NodeInterface interface {
  node()
}

type Node struct {
  Left  interface{}
  Right interface{}
}

func (node *Node) node() {}
