package managers

import (
  "github.com/chuckpreslar/codex/tree/managers"
  "github.com/chuckpreslar/codex/tree/nodes"
  "testing"
)

func TestDeleteManager(t *testing.T) {
  relation := nodes.Relation("table")
  mgr := managers.Deletion(relation)

  // The following struct members should exist.
  _ = mgr.Tree
  _ = mgr.Engine

  // The following receiver methods should exist.
  _ = mgr.Delete(1)
  _ = mgr.Engine(1)
  _, _ = mgr.ToSql()
}
