// Package nodes provides nodes to use in codex AST's.
package nodes

type AlterStatementNode struct {
  Relation    *RelationNode
  Columns     []*UnexistingColumnNode
  Constraints []*ConstraintNode
  Engine      *EngineNode
  Create      bool
}

// AlterStatementNode factory method.
func AlterStatement(relation *RelationNode, create bool) (statement *AlterStatementNode) {
  statement = new(AlterStatementNode)
  statement.Relation = relation
  statement.Columns = make([]*UnexistingColumnNode, 0)
  statement.Constraints = make([]*ConstraintNode, 0)
  statement.Create = create
  return
}
