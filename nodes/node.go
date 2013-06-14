package nodes

type Node interface {
  Left() interface{}
  Right() interface{}
  SetLeft(interface{})
  SetRight(interface{})
}

type BaseNode struct {
  left  interface{}
  right interface{}
}
