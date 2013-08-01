package managers

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

// UpdateManager manages a tree that compiles to a SQL update statement.
type UpdateManager struct {
  Tree   *nodes.UpdateStatementNode // The AST for the SQL UPDATE statement.
  engine interface{}                // The SQL Engine.
}

// Set appends to the trees Values slice a list of UnqualifiedColumnNodes
// which are to be modified in the query.
func (self *UpdateManager) Set(columns ...interface{}) *UpdateManager {
  for _, column := range columns {
    self.Tree.Values = append(self.Tree.Values, nodes.UnqualifiedColumn(column))
  }
  return self
}

// To alters the trees Values slice to be an AssignmentNode, containing the
// column from Set at the same index of the value.
func (self *UpdateManager) To(values ...interface{}) *UpdateManager {
  for index, value := range values {
    if index < len(self.Tree.Values) {
      column := self.Tree.Values[index]
      self.Tree.Values[index] = nodes.Assignment(column, value)
    }
  }
  return self
}

// Appends an expression to the current tree's Wheres slice,
// typically a comparison, i.e. 1 = 1
func (self *UpdateManager) Where(expr interface{}) *UpdateManager {
  self.Tree.Wheres = append(self.Tree.Wheres, expr)
  return self
}

// Sets the Tree's Limit to the given integer.
func (self *UpdateManager) Limit(expr interface{}) *UpdateManager {
  self.Tree.Limit = nodes.Limit(expr)
  return self
}

// Sets the SQL Enginge.
func (self *UpdateManager) Engine(engine interface{}) *UpdateManager {
  self.engine = engine
  return self
}

// Calls a visitor's Accept method based on the manager's SQL Engine.
func (self *UpdateManager) ToSql() (string, error) {
  if nil == self.engine {
    self.engine = "to_sql"
  }
  return VisitorFor(self.engine).Accept(self.Tree)
}

// UpdateManager factory method.
func Modification(relation *nodes.RelationNode) (modification *UpdateManager) {
  modification = new(UpdateManager)
  modification.Tree = nodes.UpdateStatement(relation)
  return
}
