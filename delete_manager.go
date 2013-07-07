package codex

import (
  "codex/tree/nodes"
)

type DeleteManager struct {
  Tree   *nodes.DeleteStatement
  Engine interface{}
}

func (mgmt *DeleteManager) Delete(expr interface{}) *DeleteManager {
  mgmt.Tree.Wheres = append(mgmt.Tree.Wheres, expr)
  return mgmt
}

func (mgmt *DeleteManager) ToSql() string {
  if nil == mgmt.Engine {
    mgmt.Engine = "to_sql"
  }
  return VISITORS[mgmt.Engine].Accept(mgmt.Tree)
}
