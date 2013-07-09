package nodes

// Relation is a specific Binary node
type RelationNode struct {
  Name  string // Relation's Name
  Alias string // Relation's Alias
}

func Relation(name string) *RelationNode {
  relation := new(RelationNode)
  relation.Name = name
  return relation
}
