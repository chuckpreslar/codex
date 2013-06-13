package librarian

type Comparable interface {
  compare()
}

type OrNode Node
type EqNode Node
type NeqNode Node
type GtNode Node
type GteNode Node
type LtNode Node
type LteNode Node
type MatchesNode Node

func (n OrNode) compare()      {}
func (n EqNode) compare()      {}
func (n NeqNode) compare()     {}
func (n GtNode) compare()      {}
func (n GteNode) compare()     {}
func (n LtNode) compare()      {}
func (n LteNode) compare()     {}
func (n MatchesNode) compare() {}

func (n OrNode) Or(other Comparable) OrNode {
  return OrNode{n, other}
}
func (n EqNode) Or(other Comparable) OrNode {
  return OrNode{n, other}
}
func (n NeqNode) Or(other Comparable) OrNode {
  return OrNode{n, other}
}
func (n GtNode) Or(other Comparable) OrNode {
  return OrNode{n, other}
}
func (n GteNode) Or(other Comparable) OrNode {
  return OrNode{n, other}
}
func (n LtNode) Or(other Comparable) OrNode {
  return OrNode{n, other}
}
func (n LteNode) Or(other Comparable) OrNode {
  return OrNode{n, other}
}
func (n MatchesNode) Or(other Comparable) OrNode {
  return OrNode{n, other}
}
