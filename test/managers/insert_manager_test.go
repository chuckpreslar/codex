package managers

import (
  "github.com/chuckpreslar/codex/tree/managers"
  "github.com/chuckpreslar/codex/tree/nodes"
  "testing"
)

func TestInsertManager(t *testing.T) {
  relation := nodes.Relation("table")
  mgr := managers.Insertion(relation)

  // The following struct members should exist.
  _ = mgr.Tree
  _ = mgr.Engine

  // The following receiver methods should exist.
  _ = mgr.Insert(1)
  _ = mgr.Into(1)
  _ = mgr.SetEngine(1)
  _, _ = mgr.ToSql()
}
