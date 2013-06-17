package nodes

type Node struct {
  Left  interface{}
  Right interface{}
}

func (node *Node) node() {}
