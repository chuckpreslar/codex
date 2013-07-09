package managers

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

// DeleteManager manages a tree that compiles to a SQL delete statement.
type DeleteManager struct {
  Tree   *nodes.DeleteStatementNode // The AST for the SQL DELETE statement.
  Engine interface{}                // The SQL Engine.
}

// Appends the expression to the Trees Wheres slice.
func (mgmt *DeleteManager) Delete(expr interface{}) *DeleteManager {
  mgmt.Tree.Wheres = append(mgmt.Tree.Wheres, expr)
  return mgmt
}

// Sets the SQL Enginge.
func (mgmt *DeleteManager) SetEngine(engine interface{}) *DeleteManager {
  if _, ok := VISITORS[engine]; ok {
    mgmt.Engine = engine
  }
  return mgmt
}

// Calls a visitor's Accept method based on the manager's SQL Engine.
func (mgmt *DeleteManager) ToSql() (string, error) {
  if nil == mgmt.Engine {
    mgmt.Engine = "to_sql"
  }
  return VISITORS[mgmt.Engine].Accept(mgmt.Tree)
}

// DeleteManager factory methods.
func Deletion(relation *nodes.RelationNode) *DeleteManager {
  deletion := new(DeleteManager)
  deletion.Tree = nodes.DeleteStatement(relation)
  return deletion
}
