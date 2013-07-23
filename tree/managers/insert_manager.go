package managers

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

// InsertManager manages a tree that compiles to a SQL insert statement.
type InsertManager struct {
  Tree   *nodes.InsertStatementNode // The AST for the SQL INSERT statement.
  engine interface{}                // The SQL Engine.
}

// Appends the values to the trees Values node
func (self *InsertManager) Insert(values ...interface{}) *InsertManager {
  self.Tree.Values.Expressions = append([]interface{}{}, values...)
  return self
}

// Appends the columns to the trees Columns slice and Values node.
func (self *InsertManager) Into(columns ...interface{}) *InsertManager {
  self.Tree.Values.Columns = append([]interface{}{}, columns...)
  self.Tree.Columns = append([]interface{}{}, columns...)
  return self
}

// Return sets the InsertStatementNodes Return to the `column` paramenter
// after ensureing it is a ColumnNode.
func (self *InsertManager) Returning(column interface{}) *InsertManager {
  if _, ok := column.(*nodes.ColumnNode); !ok {
    column = nodes.Column(column)
  }
  self.Tree.Returning = column
  return self
}

// Sets the SQL Enginge.
func (self *InsertManager) Engine(engine interface{}) *InsertManager {
  if _, ok := VISITORS[engine]; ok {
    self.engine = engine
  }
  return self
}

// Calls a visitor's Accept method based on the manager's SQL Engine.
func (self *InsertManager) ToSql() (string, error) {
  if nil == self.engine {
    self.engine = "to_sql"
  }
  return VISITORS[self.engine].Accept(self.Tree)
}

func Insertion(relation *nodes.RelationNode) (insertion *InsertManager) {
  insertion = new(InsertManager)
  insertion.Tree = nodes.InsertStatement(relation)
  return
}
