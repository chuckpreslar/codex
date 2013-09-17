// Package managers provides AST managers for the codex package.
package managers

import (
  "github.com/chuckpreslar/codex/nodes"
  "github.com/chuckpreslar/codex/sql"
)

// AlterManager manages a tree that compiles to SQL for create and alter statement.
type AlterManager struct {
  Tree    *nodes.AlterStatementNode // The AST for the SQL CREATE/ALTER TABLE statements.
  adapter interface{}               // The SQL adapter.
}

// AddColumn adds a UnexistingColumn from the nodes package to the AST for creation.
func (self *AlterManager) AddColumn(name interface{}, typ sql.Type) *AlterManager {
  if _, ok := name.(string); ok {
    name = nodes.UnqualifiedColumn(name)
  }

  self.Tree.UnexistingColumns = append(self.Tree.UnexistingColumns, nodes.UnexistingColumn(name, typ))
  return self
}

// AddColumn adds a UnexistingColumn from the nodes package to the AST for creation.
func (self *AlterManager) AlterColumn(name interface{}, typ sql.Type) *AlterManager {
  if _, ok := name.(string); ok {
    name = nodes.UnqualifiedColumn(name)
  }

  self.Tree.ModifiedColumns = append(self.Tree.ModifiedColumns, nodes.ExistingColumn(name, typ))
  return self
}

// AddColumn adds a ConstraintNode from the nodes package to the AST to apply to a column.
func (self *AlterManager) AddConstraint(columns []interface{}, kind sql.Constraint, options ...interface{}) *AlterManager {
  for _, column := range columns {
    if _, ok := column.(string); ok {
      column = nodes.UnqualifiedColumn(column)
    }
  }

  var node interface{}

  switch kind {
  case sql.NotNull:
    node = nodes.NotNull(columns, options...)
  case sql.Unique:
    node = nodes.Unique(columns, options...)
  case sql.PrimaryKey:
    node = nodes.PrimaryKey(columns, options...)
  case sql.ForeignKey:
    node = nodes.ForeignKey(columns, options...)
  case sql.Check:
    node = nodes.Check(columns, options...)
  case sql.Default:
    node = nodes.Default(columns, options...)
  default:
    node = nodes.Constraint(columns, options...)
  }

  self.Tree.Constraints = append(self.Tree.Constraints, node)
  return self
}

func (self *AlterManager) RemoveColumn(column interface{}) *AlterManager {
  if _, ok := column.(string); ok {
    column = nodes.UnqualifiedColumn(column)
  }

  self.Tree.RemovedColumns = append(self.Tree.RemovedColumns, column)

  return self
}

func (self *AlterManager) RemoveIndex(name interface{}) *AlterManager {
  if _, ok := name.(string); ok {
    name = nodes.IndexName(name)
  }

  self.Tree.RemovedIndicies = append(self.Tree.RemovedIndicies, name)

  return self
}

// Sets the SQL Adapter.
func (self *AlterManager) SetAdapter(adapter interface{}) *AlterManager {
  self.adapter = adapter
  return self
}

// ToSql calls a visitor's Accept method based on the manager's SQL adapter.
func (self *AlterManager) ToSql() (string, error) {
  if nil == self.adapter {
    self.adapter = "to_sql"
  }

  return VisitorFor(self.adapter).Accept(self.Tree)
}

// SelectManager factory method.
func Alteration(relation *nodes.RelationNode) (statement *AlterManager) {
  statement = new(AlterManager)
  statement.Tree = nodes.AlterStatement(relation)
  return
}
