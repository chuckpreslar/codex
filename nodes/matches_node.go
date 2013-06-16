package nodes

type MatchesNode struct {
  *Node
}

func (matches *MatchesNode) Or(other ComparisonInterface) OrInterface {
  return Or(matches, other)
}

func (matches *MatchesNode) Attribute() AttributeInterface {
  return matches.Left().(AttributeInterface)
}

func (matches *MatchesNode) Value() interface{} {
  return matches.Right()
}

func (matches *MatchesNode) Left() AttributeInterface {
  return matches.Attribute()
}

func Matches(attribute AttributeInterface, value interface{}) *MatchesNode {
  return &MatchesNode{&Node{attribute, value}}
}
