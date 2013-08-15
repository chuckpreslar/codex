package managers

import (
  "github.com/chuckpreslar/codex/managers"
  "github.com/chuckpreslar/codex/nodes"
  "github.com/chuckpreslar/codex/sql"
  "testing"
)

func TestAlterManager(t *testing.T) {
  relation := nodes.Relation("table")
  mgr := managers.Alteration(relation, true)

  // The following struct members should exist.
  _ = mgr.Tree

  // The following receiver methods should exist.
  _ = mgr.AddColumn(1, sql.STRING)
  _ = mgr.AddConstraint(1, sql.UNIQUE, 1, 2, 3)
  _ = mgr.SetEngine(1)
  _ = mgr.SetAdapter(1)
  _, _ = mgr.ToSql()
}
