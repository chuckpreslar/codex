package nodes

type RelationNode struct {
  *Node
  Name  string
  Alias string
}

func (relation *RelationNode) As(alias string) *RelationNode {
  relation.Alias = alias
  return relation
}

func Relation(name string) *RelationNode {
  return &RelationNode{&Node{nil, nil}, name, ""}
}
