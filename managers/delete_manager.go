// Package managers provides AST managers for the codex package.
package managers

import (
  "github.com/chuckpreslar/codex/nodes"
)

// DeleteManager manages a tree that compiles to a SQL delete statement.
type DeleteManager struct {
  Tree    *nodes.DeleteStatementNode // The AST for the SQL DELETE statement.
  adapter interface{}                // The SQL adapter.
}

// Appends the expression to the Trees Wheres slice.
func (self *DeleteManager) Delete(expr interface{}) *DeleteManager {
  self.Tree.Wheres = append(self.Tree.Wheres, expr)
  return self
}

// Sets the SQL Adapter.
func (self *DeleteManager) SetAdapter(adapter interface{}) *DeleteManager {
  self.adapter = adapter
  return self
}

// ToSql calls a visitor's Accept method based on the manager's SQL adapter.
func (self *DeleteManager) ToSql() (string, error) {
  if nil == self.adapter {
    self.adapter = "to_sql"
  }

  return VisitorFor(self.adapter).Accept(self.Tree)
}

// DeleteManager factory methods.
func Deletion(relation *nodes.RelationNode) (deletion *DeleteManager) {
  deletion = new(DeleteManager)
  deletion.Tree = nodes.DeleteStatement(relation)
  return
}
