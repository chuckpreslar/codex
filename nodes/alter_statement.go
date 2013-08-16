// Package nodes provides nodes to use in codex AST's.
package nodes

// AlterStatement is the base node for SQL Create and Alter Statements.
type AlterStatementNode struct {
  Relation *RelationNode // The RelationNode the AlterStatementNode affects.
  Engine   *EngineNode   // Engine the table uses if created.
  Create   bool          // Is the AlterStatementNode creating a table based on the RelationNode?
  Drop     bool          // Is the AlterStatementNode dropping a tabled based on the RelationNode?

  // FIXME: UnexistingColumns, ModifiedColumns, Constraints, RemovedColumns and
  // RemovedIndicies fields should be placed into an AlterCoreNode for visitation
  // simplification.
  UnexistingColumns []*UnexistingColumnNode // Columns to create with the alteration.
  ModifiedColumns   []*ExistingColumnNode   // Columns to modify with the alteration.
  Constraints       []interface{}           // Constraints to apply with the alteration.
  RemovedColumns    []interface{}           // Columns to be removed.
  RemovedIndicies   []interface{}           // Indicies to be removed.
}

// AlterStatementNode factory method.
func AlterStatement(relation *RelationNode, create bool) (statement *AlterStatementNode) {
  statement = new(AlterStatementNode)
  statement.Relation = relation
  statement.UnexistingColumns = make([]*UnexistingColumnNode, 0)
  statement.ModifiedColumns = make([]*ExistingColumnNode, 0)
  statement.RemovedColumns = make([]interface{}, 0)
  statement.RemovedIndicies = make([]interface{}, 0)
  statement.Constraints = make([]interface{}, 0)
  statement.Create = create
  return
}
