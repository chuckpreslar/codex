package managers

import (
  "github.com/chuckpreslar/codex/managers"
  "github.com/chuckpreslar/codex/nodes"
  "testing"
)

func TestInsertManager(t *testing.T) {
  relation := nodes.Relation("table")
  mgr := managers.Insertion(relation)

  // The following struct members should exist.
  _ = mgr.Tree

  // The following receiver methods should exist.
  _ = mgr.Insert(1)
  _ = mgr.Into(1)
  _ = mgr.Returning(1)
  _ = mgr.SetAdapter(1)
  _, _ = mgr.ToSql()
}
