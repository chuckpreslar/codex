package managers

import (
  "github.com/chuckpreslar/codex/managers"
  "github.com/chuckpreslar/codex/nodes"
  "testing"
)

func TestUpdateManager(t *testing.T) {
  relation := nodes.Relation("table")
  mgr := managers.Modification(relation)

  // The following struct members should exist.
  _ = mgr.Tree

  // The following receiver methods should exist.
  _ = mgr.Set(1)
  _ = mgr.To(1)
  _ = mgr.Where(1)
  _ = mgr.Limit(1)
  _ = mgr.SetAdapter(1)
  _, _ = mgr.ToSql()
}
