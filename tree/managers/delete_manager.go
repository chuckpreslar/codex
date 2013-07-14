package managers

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

// DeleteManager manages a tree that compiles to a SQL delete statement.
type DeleteManager struct {
  Tree   *nodes.DeleteStatementNode // The AST for the SQL DELETE statement.
  engine interface{}                // The SQL Engine.
}

// Appends the expression to the Trees Wheres slice.
func (self *DeleteManager) Delete(expr interface{}) *DeleteManager {
  self.Tree.Wheres = append(self.Tree.Wheres, expr)
  return self
}

// Sets the SQL Enginge.
func (self *DeleteManager) Engine(engine interface{}) *DeleteManager {
  if _, ok := VISITORS[engine]; ok {
    self.engine = engine
  }
  return self
}

// Calls a visitor's Accept method based on the manager's SQL Engine.
func (self *DeleteManager) ToSql() (string, error) {
  if nil == self.engine {
    self.engine = "to_sql"
  }
  return VISITORS[self.engine].Accept(self.Tree)
}

// DeleteManager factory methods.
func Deletion(relation *nodes.RelationNode) (deletion *DeleteManager) {
  deletion = new(DeleteManager)
  deletion.Tree = nodes.DeleteStatement(relation)
  return
}
