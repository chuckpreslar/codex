package nodes

type Node interface {
  Left() interface{}
  Right() interface{}
  SetLeft(interface{})
  SetRight(interface{})
}

type Comparison interface {
  Node
  compare()
}

type Projection interface {
  Node
  project()
}

type BaseNode struct {
  left  interface{}
  right interface{}
}

func (n BaseNode) Left() interface{} {
  return n.left
}

func (n BaseNode) Right() interface{} {
  return n.right
}

func (n BaseNode) SetLeft(value interface{}) {
  n.left = value
}

func (n BaseNode) SetRight(value interface{}) {
  n.right = value
}
