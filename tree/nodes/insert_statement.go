package nodes

// InsertStatement is the base node for SQL Insert Statements.
type InsertStatementNode struct {
  Relation *RelationNode // Pointer to the Relation the Insert Statement is acting on.
  Columns  []interface{} // Columns the Insert Statement is effecting.
  Values   *ValuesNode   // Pointer to the Values for insertion.
}

// InsertStatementNode factory method.
func InsertStatement(relation *RelationNode) *InsertStatementNode {
  stmt := new(InsertStatementNode)
  stmt.Relation = relation
  stmt.Columns = make([]interface{}, 0)
  stmt.Values = Values()
  return stmt
}
