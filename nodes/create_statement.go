// Package nodes provides nodes to use in codex AST's.
package nodes

type CreateStatementNode struct {
  Relation *RelationNode
  Columns []*UnexistingColumnNode
  Engine  *EngineNode
}

func CreateStatement(relation *RelationNode) (statement *CreateStatementNode) {
  statement = new(CreateStatementNode)
  statement.Relation = relation
  statement.Columns = make([]*UnexistingColumnNode, 0)
  return
}
