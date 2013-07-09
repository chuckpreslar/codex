package managers

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

type InsertManager struct {
  Tree   *nodes.InsertStatementNode
  Engine interface{}
}

func (mgmt *InsertManager) Insert(values ...interface{}) *InsertManager {
  mgmt.Tree.Values.Expressions = append([]interface{}{}, values...)

  return mgmt
}

func (mgmt *InsertManager) Into(columns ...interface{}) *InsertManager {
  mgmt.Tree.Values.Columns = append([]interface{}{}, columns...)
  mgmt.Tree.Columns = append([]interface{}{}, columns...)

  return mgmt
}

func (mgmt *InsertManager) SetEngine(engine interface{}) *InsertManager {
  if _, ok := VISITORS[engine]; ok {
    mgmt.Engine = engine
  }

  return mgmt
}

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
