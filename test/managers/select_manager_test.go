package managers

import (
  "github.com/chuckpreslar/codex/tree/managers"
  "github.com/chuckpreslar/codex/tree/nodes"
  "testing"
)

func TestSelectManager(t *testing.T) {
  relation := nodes.Relation("table")
  mgr := managers.Selection(relation)

  // The following struct members should exist.
  _ = mgr.Tree
  _ = mgr.Context
  _ = mgr.Engine

  // The following receiver methods should exist.
  _ = mgr.Project(1)
  _ = mgr.Where(1)
  _ = mgr.Offset(1)
  _ = mgr.Limit(1)
  _ = mgr.InnerJoin(1)
  _ = mgr.OuterJoin(1)
  _ = mgr.On(1)
  _ = mgr.Order(1)
  _ = mgr.Group(1)
  _ = mgr.Having(1)
  _ = mgr.Union(managers.Selection(relation))
  _ = mgr.Engine(1)
  _, _ = mgr.ToSql()
}
