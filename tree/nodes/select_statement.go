package nodes

// SelectStatement is the base node for SQL Select Statements.
type SelectStatementNode struct {
  Cores  []*SelectCoreNode // An array of SelectCores.
  Limit  *LimitNode        // Potential Limit node for limiting the number of results returned.
  Offset *OffsetNode       // Potential Offset node for skipping records.
}

func SelectStatement(relation *RelationNode) *SelectStatementNode {
  stmt := new(SelectStatementNode)
  stmt.Cores = []*SelectCoreNode{SelectCore(relation)}
  return stmt
}
