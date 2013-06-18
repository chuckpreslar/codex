package nodes

type OrNode struct {
  *Node
}

func (or *OrNode) Or(b interface{}) *OrNode {
  return Or(or, b)
}

func (or *OrNode) And(b interface{}) *AndNode {
  return And(or, b)
}

func Or(a, b interface{}) *OrNode {
  return &OrNode{&Node{a, b}}
}
