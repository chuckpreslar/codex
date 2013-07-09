package nodes

// DeleteStatement is the base node for SQL Delete Statements.
type DeleteStatementNode struct {
  Relation *RelationNode // Pointer to the Relation the Delete Statement is acting on.
  Wheres   []interface{} // Wheres is an array of expressions/nodes.
}

// DeleteStatementNode factory method.
func DeleteStatement(relation *RelationNode) *DeleteStatementNode {
  stmt := new(DeleteStatementNode)
  stmt.Relation = relation
  return stmt
}
