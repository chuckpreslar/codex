package nodes

type DoesNotMatchNode struct {
  *Node
}

func DoesNotMatch(a, b interface{}) *ComparisonNode {
  return Comparison(&DoesNotMatchNode{&Node{a, b}})
}
