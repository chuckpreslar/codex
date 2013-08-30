// Package managers provides AST managers for the codex package.
package managers

import (
  "github.com/chuckpreslar/codex/nodes"
  "github.com/chuckpreslar/codex/sql"
)

type CreateManager struct {
  Tree    *nodes.CreateStatementNode
  adapter interface{} // The SQL adapter.
}

// AddColumn adds a UnexistingColumn from the nodes package to the AST for creation.
func (self *CreateManager) AddColumn(name interface{}, typ sql.Type) *CreateManager {
  if _, ok := name.(string); ok {
    name = nodes.UnqualifiedColumn(name)
  }

  self.Tree.UnexistingColumns = append(self.Tree.UnexistingColumns, nodes.UnexistingColumn(name, typ))
  return self
}

// AddColumn adds a ConstraintNode from the nodes package to the AST to apply to a column.
func (self *CreateManager) AddConstraint(column interface{}, kind sql.Constraint, options ...interface{}) *CreateManager {
  if _, ok := column.(string); ok {
    column = nodes.UnqualifiedColumn(column)
  }

  var node interface{}

  switch kind {
  case sql.NOT_NULL:
    node = nodes.NotNull(column, options...)
  case sql.UNIQUE:
    node = nodes.Unique(column, options...)
  case sql.PRIMARY_KEY:
    node = nodes.PrimaryKey(column, options...)
  case sql.FOREIGN_KEY:
    node = nodes.ForeignKey(column, options...)
  case sql.CHECK:
    node = nodes.Check(column, options...)
  case sql.DEFAULT:
    node = nodes.Default(column, options...)
  default:
    node = nodes.Constraint(column, options...)
  }

  self.Tree.Constraints = append(self.Tree.Constraints, node)
  return self
}

// SetEngine sets the AST's Engine field, used for table creation.
func (self *CreateManager) SetEngine(engine interface{}) *CreateManager {
  if _, ok := engine.(*nodes.EngineNode); !ok {
    engine = nodes.Engine(engine)
  }

  self.Tree.Engine = engine.(*nodes.EngineNode)
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

// SelectManager factory method.
func Creation(relation *nodes.RelationNode) (statement *CreateManager) {
  statement = new(CreateManager)
  statement.Tree = nodes.CreateStatement(relation)
  return
}
