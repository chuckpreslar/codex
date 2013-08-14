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
  if _, ok := name.(string); ok {
    name = nodes.UnqualifiedColumn(name)
  }

  self.Tree.Columns = append(self.Tree.Columns, nodes.UnexistingColumn(name, typ))
  return self
}

func (self *AlterManager) AddConstraint(column interface{}, kind sql.Constraint, options ...interface{}) *AlterManager {
  if _, ok := column.(string); ok {
    column = nodes.UnqualifiedColumn(column)
  }

  self.Tree.Constraints = append(self.Tree.Constraints, nodes.Constraint(column, kind, options...))
  return self
}

func (self *AlterManager) SetEngine(engine interface{}) *AlterManager {
  if _, ok := engine.(*nodes.EngineNode); !ok {
    engine = nodes.Engine(engine)
  }

  self.Tree.Engine = engine.(*nodes.EngineNode)
  return self
}

func (self *AlterManager) SetAdapter(adapter interface{}) *AlterManager {
  self.adapter = adapter
  return self
}

func (self *AlterManager) ToSql() (string, error) {
  if nil == self.adapter {
    self.adapter = "to_sql"
  }

  return VisitorFor(self.adapter).Accept(self.Tree)
}

// SelectManager factory method.
func Alteration(relation *nodes.RelationNode, create bool) (statement *AlterManager) {
  statement = new(AlterManager)
  statement.Tree = nodes.AlterStatement(relation, create)
  return
}
