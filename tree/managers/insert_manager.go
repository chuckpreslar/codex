package managers

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

// InsertManager manages a tree that compiles to a SQL insert statement.
type InsertManager struct {
  Tree   *nodes.InsertStatementNode // The AST for the SQL INSERT statement.
  Engine interface{}                // The SQL Engine.
}

// Appends the values to the trees Values node
func (mgmt *InsertManager) Insert(values ...interface{}) *InsertManager {
  mgmt.Tree.Values.Expressions = append([]interface{}{}, values...)
  return mgmt
}

// Appends the columns to the trees Columns slice and Values node.
func (mgmt *InsertManager) Into(columns ...interface{}) *InsertManager {
  mgmt.Tree.Values.Columns = append([]interface{}{}, columns...)
  mgmt.Tree.Columns = append([]interface{}{}, columns...)
  return mgmt
}

// Sets the SQL Enginge.
func (mgmt *InsertManager) SetEngine(engine interface{}) *InsertManager {
  if _, ok := VISITORS[engine]; ok {
    mgmt.Engine = engine
  }
  return mgmt
}

// Calls a visitor's Accept method based on the manager's SQL Engine.
func (mgmt *InsertManager) ToSql() (string, error) {
  if nil == mgmt.Engine {
    mgmt.Engine = "to_sql"
  }
  return VISITORS[mgmt.Engine].Accept(mgmt.Tree)
}

func Insertion(relation *nodes.RelationNode) *InsertManager {
  insertion := new(InsertManager)
  insertion.Tree = nodes.InsertStatement(relation)
  return insertion
}
