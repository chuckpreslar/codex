package nodes

type RelationNode struct {
  *Node
  aliases []string
}

func (relation *RelationNode) Name() string {
  return relation.Left.(string)
}

func (relation *RelationNode) Aliases() []string {
  return relation.aliases
}

func (relation *RelationNode) AddAliases(aliases ...string) RelationInterface {
  relation.aliases = append(relation.aliases, aliases...)
  return relation
}

func Relation(name string, aliases ...string) *RelationNode {
  return &RelationNode{&Node{name, nil}, aliases}
}
