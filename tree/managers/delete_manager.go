package managers

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

type DeleteManager struct {
  Tree   *nodes.DeleteStatementNode
  Engine interface{}
}

func (mgmt *DeleteManager) Delete(expr interface{}) *DeleteManager {
  mgmt.Tree.Wheres = append(mgmt.Tree.Wheres, expr)
  return mgmt
}

func (mgmt *DeleteManager) SetEngine(engine interface{}) *DeleteManager {
  if _, ok := VISITORS[engine]; ok {
    mgmt.Engine = engine
  }

  return mgmt
}

func (mgmt *DeleteManager) ToSql() (string, error) {
  if nil == mgmt.Engine {
    mgmt.Engine = "to_sql"
  }

  return VISITORS[mgmt.Engine].Accept(mgmt.Tree)
}

func Deletion(relation *nodes.RelationNode) *DeleteManager {
  deletion := new(DeleteManager)
  deletion.Tree = nodes.DeleteStatement(relation)
  return deletion
}
