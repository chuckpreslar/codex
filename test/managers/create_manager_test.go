package managers

import (
  "github.com/chuckpreslar/codex/managers"
  "github.com/chuckpreslar/codex/nodes"
  "github.com/chuckpreslar/codex/sql"
  "testing"
)

func TestCreateManager(t *testing.T) {
  relation := nodes.Relation("table")
  mgr := managers.Creation(relation)

  // The following struct members should exist.
  _ = mgr.Tree

  // The following receiver methods should exist.
  _ = mgr.AddColumn(1, sql.STRING)
  _ = mgr.SetEngine(1)
  _ = mgr.SetAdapter(1)
  _, _ = mgr.ToSql()
}
