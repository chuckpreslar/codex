package nodes

// InsertStatement is the base node for SQL Insert Statements.
type InsertStatementNode struct {
  Relation  *RelationNode // Pointer to the Relation the Insert Statement is acting on.
  Columns   []interface{} // Columns the Insert Statement is effecting.
  Returning interface{}   // Columns to return after the Insert Statement is executed.
  Values    *ValuesNode   // Pointer to the Values for insertion.
}

// InsertStatementNode factory method.
func InsertStatement(relation *RelationNode) (statement *InsertStatementNode) {
  statement = new(InsertStatementNode)
  statement.Relation = relation
  statement.Columns = make([]interface{}, 0)
  statement.Values = Values()
  return
}
