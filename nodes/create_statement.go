// Package nodes provides nodes to use in codex AST's.
package nodes

type CreateStatementNode struct {
	Relation          *RelationNode           //
	Engine            *EngineNode             // Engine the table uses if created.
	UnexistingColumns []*UnexistingColumnNode // Columns to create with the alteration.
	Constraints       []interface{}           // Constraints to apply with the alteration.
}

func CreateStatement(relation *RelationNode) (statement *CreateStatementNode) {
	statement = new(CreateStatementNode)
	statement.Relation = relation
	statement.UnexistingColumns = make([]*UnexistingColumnNode, 0)
	statement.Constraints = make([]interface{}, 0)
	return
}
