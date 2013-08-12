// Package nodes provides nodes to use in codex AST's.
package nodes

// SelectStatement is the base node for SQL Select Statements.
type SelectStatementNode struct {
  Cores      []*SelectCoreNode // An array of SelectCores.
  Orders     []interface{}     // An array of nodes for ordering results.
  Combinator interface{}       // Potential Union/Intersect/Except node.
  Limit      *LimitNode        // Potential Limit node for limiting the number of results returned.
  Offset     *OffsetNode       // Potential Offset node for skipping records.
}

// SelectStatementNode factory method.
func SelectStatement(relation *RelationNode) (statement *SelectStatementNode) {
  statement = new(SelectStatementNode)
  statement.Orders = make([]interface{}, 0)
  statement.Cores = []*SelectCoreNode{SelectCore(relation)}
  return
}
