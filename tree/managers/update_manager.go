package managers

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

// UpdateManager manages a tree that compiles to a SQL update statement.
type UpdateManager struct {
  Tree   *nodes.UpdateStatementNode // The AST for the SQL UPDATE statement.
  Engine interface{}                // The SQL Engine.
}

// Set appends to the trees Values slice a list of UnqualifiedColumnNodes
// which are to be modified in the query.
func (mgmt *UpdateManager) Set(columns ...interface{}) *UpdateManager {
  for _, column := range columns {
    mgmt.Tree.Values = append(mgmt.Tree.Values, nodes.UnqualifiedColumn(column))
  }
  return mgmt
}

// To alters the trees Values slice to be an AssignmentNode, containing the
// column from Set at the same index of the value.
func (mgmt *UpdateManager) To(values ...interface{}) *UpdateManager {
  for index, value := range values {
    if index < len(mgmt.Tree.Values) {
      column := mgmt.Tree.Values[index]
      mgmt.Tree.Values[index] = nodes.Assignment(column, value)
    }
  }
  return mgmt
}

// Appends an expression to the current tree's Wheres slice,
// typically a comparison, i.e. 1 = 1
func (mgmt *UpdateManager) Where(expr interface{}) *UpdateManager {
  mgmt.Tree.Wheres = append(mgmt.Tree.Wheres, expr)
  return mgmt
}

// Sets the Tree's Limit to the given integer.
func (mgmt *UpdateManager) Limit(expr interface{}) *UpdateManager {
  mgmt.Tree.Limit = nodes.Limit(expr)
  return mgmt
}

// Sets the SQL Enginge.
func (mgmt *UpdateManager) SetEngine(engine interface{}) *UpdateManager {
  if _, ok := VISITORS[engine]; ok {
    mgmt.Engine = engine
  }
  return mgmt
}

// Calls a visitor's Accept method based on the manager's SQL Engine.
func (mgmt *UpdateManager) ToSql() (string, error) {
  if nil == mgmt.Engine {
    mgmt.Engine = "to_sql"
  }
  return VISITORS[mgmt.Engine].Accept(mgmt.Tree)
}

// UpdateManager factory method.
func Modification(relation *nodes.RelationNode) *UpdateManager {
  modification := new(UpdateManager)
  modification.Tree = nodes.UpdateStatement(relation)
  return modification
}
