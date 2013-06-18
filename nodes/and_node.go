package nodes

type AndNode struct {
  *Node
}

func (and *AndNode) Or(b interface{}) *OrNode {
  return Or(and, b)
}

func (and *AndNode) And(b interface{}) *AndNode {
  return And(and, b)
}

func And(a, b interface{}) *AndNode {
  return &AndNode{&Node{a, b}}
}
