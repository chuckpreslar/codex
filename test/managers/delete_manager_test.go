package managers

import (
  "github.com/chuckpreslar/codex/managers"
  "github.com/chuckpreslar/codex/nodes"
  "testing"
)

func TestDeleteManager(t *testing.T) {
  relation := nodes.Relation("table")
  mgr := managers.Deletion(relation)

  // The following struct members should exist.
  _ = mgr.Tree

  // The following receiver methods should exist.
  _ = mgr.Delete(1)
  _ = mgr.SetAdapter(1)
  _, _ = mgr.ToSql()
}
