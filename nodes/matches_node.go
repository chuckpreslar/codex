package nodes

type MatchesNode struct {
  *Node
}

func (matches *MatchesNode) Or(other ComparisonInterface) ComparisonInterface {
  return Or(matches, other)
}

func (matches *MatchesNode) Attribute() AttributeInterface {
  return matches.Left.(AttributeInterface)
}

func (matches *MatchesNode) Value() interface{} {
  return matches.Right
}

func Matches(value ComparableInterface, other interface{}) *MatchesNode {
  return &MatchesNode{&Node{value, other}}
}
