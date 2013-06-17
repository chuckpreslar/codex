package librarian

import (
  "librarian/nodes"
  "librarian/visitors"
)

type SelectManager struct {
  Tree    *nodes.SelectStatementNode
  Context *nodes.SelectCoreNode
}

func (mgr *SelectManager) Project(attribute nodes.AttributeInterface) *SelectManager {
  mgr.Context.AppendToProjections(attribute)
  return mgr
}

func (mgr *SelectManager) ToSql() string {
  visitor := visitors.ToSql()
  return visitor.Accept(mgr.Tree)
}

func Select(relation nodes.RelationInterface) *SelectManager {
  tree := nodes.SelectStatement(relation)
  core := tree.CoreAtIndex(0)
  return &SelectManager{tree, core}
}
