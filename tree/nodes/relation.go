package nodes

// RelationNode is a specific BinaryNode
type RelationNode struct {
  Name  string // Relation's Name
  Alias string // Relation's Alias
}

// RelationNode factory method.
func Relation(name string) (relation *RelationNode) {
  relation = new(RelationNode)
  relation.Name = name
  return
}
