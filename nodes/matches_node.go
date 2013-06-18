package nodes

type MatchesNode struct {
  *Node
}

func Matches(a, b interface{}) *ComparisonNode {
  return Comparison(&MatchesNode{&Node{a, b}})
}
