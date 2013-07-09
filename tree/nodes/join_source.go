package nodes

// JoinSource is a specific Binary node.
type JoinSourceNode struct {
  Left  *RelationNode // Left child of the JoinSource node, a pointer to a Relation.
  Right []interface{} // Right child of the JoinSource node contains joins and their instructions
}

func JoinSource(relation *RelationNode) *JoinSourceNode {
  source := new(JoinSourceNode)
  source.Left = relation
  source.Right = make([]interface{}, 0)
  return source
}
