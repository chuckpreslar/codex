// Package managers provides AST managers for the codex package.
package managers

import (
  "github.com/chuckpreslar/codex/nodes"
  "github.com/chuckpreslar/codex/sql"
)

type CreateManager struct {
  Tree    *nodes.CreateStatementNode
  adapter interface{}
}

func (self *CreateManager) AddColumn(name interface{}, typ sql.Type, constraints ...sql.Constraint) *CreateManager {
  self.Tree.Columns = append(self.Tree.Columns, nodes.UnexistingColumn(name, typ, constraints...))
  return self
}

func (self *CreateManager) SetEngine(engine interface{}) *CreateManager {
  self.Tree.Engine = nodes.Engine(engine)
  return self
}

// Sets the SQL Adapter.
func (self *CreateManager) SetAdapter(adapter interface{}) *CreateManager {
  self.adapter = adapter
  return self
}

// ToSql calls a visitor's Accept method based on the manager's SQL adapter.
func (self *CreateManager) ToSql() (string, error) {
  if nil == self.adapter {
    self.adapter = "to_sql"
  }
  return VisitorFor(self.adapter).Accept(self.Tree)
}

func Creation(relation *nodes.RelationNode) (creation *CreateManager) {
  creation = new(CreateManager)
  creation.Tree = nodes.CreateStatement(relation)
  return
}
