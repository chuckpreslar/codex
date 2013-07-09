package managers

import (
  "github.com/chuckpreslar/codex/tree/managers"
  "github.com/chuckpreslar/codex/tree/nodes"
  "testing"
)

func TestUpdateManager(t *testing.T) {
  relation := nodes.Relation("table")
  mgr := managers.Modification(relation)

  // The following struct members should exist.
  _ = mgr.Tree
  _ = mgr.Engine

  // The following receiver methods should exist.
  _ = mgr.Set(1)
  _ = mgr.To(1)
  _ = mgr.Where(1)
  _ = mgr.Limit(1)
  _ = mgr.SetEngine(1)
  _, _ = mgr.ToSql()
}
