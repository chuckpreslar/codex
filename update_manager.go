package codex

import (
  "codex/tree/nodes"
)

type UpdateManager struct {
  Tree   *nodes.UpdateStatement
  Engine interface{}
}

func (mgmt *UpdateManager) Set(columns ...interface{}) *UpdateManager {
  col := columns[0]
  switch col.(type) {
  case Values:
    return mgmt.InsertAssignments(col.(Values))
  default:
    mgmt.Tree.Values = []interface{}{}
    for _, column := range columns {
      mgmt.Tree.Values = append(mgmt.Tree.Values, &nodes.UnqualifiedColumn{column})
    }
  }
  return mgmt
}

func (mgmt *UpdateManager) InsertAssignments(values Values) *UpdateManager {
  assignments := []interface{}{}
  for column, value := range values {
    assignments = append(assignments, &nodes.Assignment{&nodes.UnqualifiedColumn{column}, value})
  }

  mgmt.Tree.Values = assignments
  return mgmt
}

func (mgmt *UpdateManager) To(values ...interface{}) *UpdateManager {
  for index, value := range values {
    if index < len(mgmt.Tree.Values) {
      column := mgmt.Tree.Values[index]
      mgmt.Tree.Values[index] = &nodes.Assignment{column, value}
    }
  }
  return mgmt
}

func (mgmt *UpdateManager) Where(expr interface{}) *UpdateManager {
  mgmt.Tree.Wheres = append(mgmt.Tree.Wheres, expr)
  return mgmt
}

func (mgmt *UpdateManager) Limit(expr interface{}) *UpdateManager {
  if nil == mgmt.Tree.Limit {
    mgmt.Tree.Limit = &nodes.Limit{}
  }

  mgmt.Tree.Limit.Expr = expr
  return mgmt
}

func (mgmt *UpdateManager) ToSql() string {
  if nil == mgmt.Engine {
    mgmt.Engine = "to_sql"
  }
  return VISITORS[mgmt.Engine].Accept(mgmt.Tree)
}
