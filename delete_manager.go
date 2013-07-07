package codex

import (
  "codex/tree/nodes"
  "codex/tree/visitors"
)

type DeleteManager struct {
  Tree *nodes.DeleteStatement
}

func (mgmt *DeleteManager) Delete(expr interface{}) *DeleteManager {
  mgmt.Tree.Wheres = append(mgmt.Tree.Wheres, expr)
  return mgmt
}

func (mgmt *DeleteManager) ToSql() string {
  visitor := &visitors.ToSqlVisitor{}
  return visitor.Accept(mgmt.Tree)
}
