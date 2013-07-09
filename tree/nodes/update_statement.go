package nodes

// UpdateStatement is the base node for SQL Update Statements.
type UpdateStatementNode struct {
  Relation *RelationNode // Pointer to the Relation the Delete Statement is acting on.
  Values   []interface{} // Values is an array of expressions/nodes.
  Wheres   []interface{} // Wheres is an array of expressions/nodes.
  Limit    *LimitNode    // Potential Limit node for limiting the number of rows effected.
}

// UpdateStatementNode factory method.
func UpdateStatement(relation *RelationNode) *UpdateStatementNode {
  stmt := new(UpdateStatementNode)
  stmt.Relation = relation
  stmt.Values = make([]interface{}, 0)
  stmt.Wheres = make([]interface{}, 0)
  return stmt
}
