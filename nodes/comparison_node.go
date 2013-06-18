package nodes

type ComparisonNode struct {
  *Node
}

func (comparison *ComparisonNode) Or(a interface{}) *OrNode {
  return Or(comparison, a)
}

func (comparison *ComparisonNode) And(a interface{}) *AndNode {
  return And(comparison, a)
}

func Comparison(a interface{}) *ComparisonNode {
  return &ComparisonNode{&Node{a, nil}}
}
