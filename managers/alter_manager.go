// Package managers provides AST managers for the codex package.
package managers

import (
  "github.com/chuckpreslar/codex/nodes"
  "github.com/chuckpreslar/codex/sql"
)

type AlterManager struct {
  Tree    *nodes.AlterStatementNode // The AST for the SQL CREATE/ALTER TABLE statements.
  adapter interface{}               // The SQL adapter.
}

func (self *AlterManager) AddColumn(name interface{}, typ sql.Type) *AlterManager {
  self.Tree.Columns = append(self.Tree.Columns, nodes.UnexistingColumn(name, typ))
  return self
}

func (self *AlterManager) AddConstraint() *AlterManager {
  return self
}

func (self *AlterManager) SetAdapter(adapter interface{}) *AlterManager {
  self.adapter = adapter
  return self
}

// SelectManager factory method.
func Alteration(relation *nodes.RelationNode, create bool) (statement *AlterManager) {
  statement = new(AlterManager)
  statement.Tree = nodes.AlterStatement(relation, create)
  return
}
