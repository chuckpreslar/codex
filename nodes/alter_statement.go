// Package nodes provides nodes to use in codex AST's.
package nodes

// AlterStatement is the base node for SQL Create and Alter Statements.
type AlterStatementNode struct {
  Relation    *RelationNode           // The RelationNode the AlterStatementNode affects.
  Columns     []*UnexistingColumnNode // Columns to create with the alteration.
  Constraints []interface{}           // Constraints to apply with the alteration.
  Engine      *EngineNode             // Engine the table uses if created.
  Create      bool                    // Is the AlterStatementNode creating a table based on the RelationNode.
}

// AlterStatementNode factory method.
func AlterStatement(relation *RelationNode, create bool) (statement *AlterStatementNode) {
  statement = new(AlterStatementNode)
  statement.Relation = relation
  statement.Columns = make([]*UnexistingColumnNode, 0)
  statement.Constraints = make([]interface{}, 0)
  statement.Create = create
  return
}
