package nodes

// UpdateStatement is the base node for SQL Update Statements.
type UpdateStatementNode struct {
  Relation *RelationNode // Pointer to the Relation the Delete Statement is acting on.
  Values   []interface{} // Values is an array of expressions/nodes.
  Wheres   []interface{} // Wheres is an array of expressions/nodes.
  Limit    *LimitNode    // Potential Limit node for limiting the number of rows effected.
}

// UpdateStatementNode factory method.
func UpdateStatement(relation *RelationNode) (statement *UpdateStatementNode) {
  statement = new(UpdateStatementNode)
  statement.Relation = relation
  statement.Values = make([]interface{}, 0)
  statement.Wheres = make([]interface{}, 0)
  return
}
