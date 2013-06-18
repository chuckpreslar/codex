package nodes

type MatchesNode struct {
  *Node
}

func Matches(a, b interface{}) *MatchesNode {
  return &MatchesNode{&Node{a, b}}
}
