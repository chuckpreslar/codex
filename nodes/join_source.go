// Package nodes provides nodes to use in codex AST's.
package nodes

// JoinSourceNode is a specific BinaryNode.
type JoinSourceNode struct {
	Left  *RelationNode // Left child of the JoinSource node, a pointer to a Relation.
	Right []interface{} // Right child of the JoinSource node contains joins and their instructions
}

// JoinSourceNode factory method.
func JoinSource(relation *RelationNode) (source *JoinSourceNode) {
	source = new(JoinSourceNode)
	source.Left = relation
	source.Right = make([]interface{}, 0)
	return
}
