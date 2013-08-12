// Package managers provides AST managers for the codex package.
package managers

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

// InsertManager manages a tree that compiles to a SQL insert statement.
type InsertManager struct {
  Tree    *nodes.InsertStatementNode // The AST for the SQL INSERT statement.
  adapter interface{}                // The SQL adapter.
}

// Appends the values to the trees Values node
func (self *InsertManager) Insert(values ...interface{}) *InsertManager {
  self.Tree.Values.Expressions = append(self.Tree.Values.Expressions, values...)
  return self
}

// Appends the columns to the trees Columns slice and Values node.
func (self *InsertManager) Into(columns ...interface{}) *InsertManager {
  self.Tree.Values.Columns = append(self.Tree.Values.Columns, columns...)
  self.Tree.Columns = append(self.Tree.Columns, columns...)
  return self
}

// Return sets the InsertStatementNodes Return to the `column` paramenter
// after ensureing it is a ColumnNode.
func (self *InsertManager) Returning(column interface{}) *InsertManager {
  if _, ok := column.(string); ok {
    column = nodes.Column(column)
  }
  self.Tree.Returning = column
  return self
}

// Sets the SQL Adapter.
func (self *InsertManager) SetAdapter(adapter interface{}) *InsertManager {
  self.adapter = adapter
  return self
}

// ToSql calls a visitor's Accept method based on the manager's SQL adapter.
func (self *InsertManager) ToSql() (string, error) {
  if nil == self.adapter {
    self.adapter = "to_sql"
  }
  return VisitorFor(self.adapter).Accept(self.Tree)
}

func Insertion(relation *nodes.RelationNode) (insertion *InsertManager) {
  insertion = new(InsertManager)
  insertion.Tree = nodes.InsertStatement(relation)
  return
}
