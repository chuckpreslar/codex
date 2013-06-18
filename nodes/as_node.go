package nodes

type AsNode struct {
  *Node
}

func (as *AsNode) Eq(other interface{}) ComparisonInterface {
  return Eq(as, other)
}
func (as *AsNode) Neq(other interface{}) ComparisonInterface {
  return Neq(as, other)
}
func (as *AsNode) Gt(other interface{}) ComparisonInterface {
  return Gt(as, other)
}
func (as *AsNode) Gte(other interface{}) ComparisonInterface {
  return Gte(as, other)
}
func (as *AsNode) Lt(other interface{}) ComparisonInterface {
  return Lt(as, other)
}
func (as *AsNode) Lte(other interface{}) ComparisonInterface {
  return Lte(as, other)
}
func (as *AsNode) Matches(other interface{}) ComparisonInterface {
  return Matches(as, other)
}
func (as *AsNode) As(other interface{}) ComparableInterface {
  return As(as, other)
}

func As(comparable ComparableInterface, alias interface{}) *AsNode {
  return &AsNode{&Node{comparable, alias}}
}
