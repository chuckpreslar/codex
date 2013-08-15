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

  self.Tree.Columns = append(self.Tree.Columns, nodes.UnexistingColumn(name, typ))
  return self
}

// AddColumn adds a ConstraintNode from the nodes package to the AST to apply to a column.
func (self *AlterManager) AddConstraint(column interface{}, kind sql.Constraint, options ...interface{}) *AlterManager {
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
func (self *AlterManager) SetEngine(engine interface{}) *AlterManager {
  if _, ok := engine.(*nodes.EngineNode); !ok {
    engine = nodes.Engine(engine)
  }

  self.Tree.Engine = engine.(*nodes.EngineNode)
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
func Alteration(relation *nodes.RelationNode, create bool) (statement *AlterManager) {
  statement = new(AlterManager)
  statement.Tree = nodes.AlterStatement(relation, create)
  return
}
