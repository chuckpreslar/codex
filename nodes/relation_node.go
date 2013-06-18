package nodes

type RelationNode struct {
  *Node
  name  string
  alias string
}

func (relation *RelationNode) As(alias string) *RelationNode {
  relation.alias = alias
  return relation
}

func Relation(name string) *RelationNode {
  return &RelationNode{&Node{nil, nil}, name, ""}
}
