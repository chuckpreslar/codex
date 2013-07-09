package nodes

// JoinSourceNode is a specific BinaryNode.
type JoinSourceNode struct {
  Left  *RelationNode // Left child of the JoinSource node, a pointer to a Relation.
  Right []interface{} // Right child of the JoinSource node contains joins and their instructions
}

// JoinSourceNode factory method.
func JoinSource(relation *RelationNode) *JoinSourceNode {
  source := new(JoinSourceNode)
  source.Left = relation
  source.Right = make([]interface{}, 0)
  return source
}
