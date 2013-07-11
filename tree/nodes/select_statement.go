package nodes

// SelectStatement is the base node for SQL Select Statements.
type SelectStatementNode struct {
  Cores  []*SelectCoreNode // An array of SelectCores.
  Orders []interface{}     // An array of nodes for ordering results.
  Limit  *LimitNode        // Potential Limit node for limiting the number of results returned.
  Offset *OffsetNode       // Potential Offset node for skipping records.
}

// SelectStatementNode factory method.
func SelectStatement(relation *RelationNode) *SelectStatementNode {
  stmt := new(SelectStatementNode)
  stmt.Orders = make([]interface{}, 0)
  stmt.Cores = []*SelectCoreNode{SelectCore(relation)}
  return stmt
}
