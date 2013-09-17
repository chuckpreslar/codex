package managers

import (
  "github.com/chuckpreslar/codex/managers"
  "github.com/chuckpreslar/codex/nodes"
  "github.com/chuckpreslar/codex/sql"
  "testing"
)

func TestAlterManager(t *testing.T) {
  relation := nodes.Relation("table")
  mgr := managers.Alteration(relation)

  // The following struct members should exist.
  _ = mgr.Tree

  // The following receiver methods should exist.
  _ = mgr.AddColumn(1, sql.String)
  _ = mgr.AddConstraint([]interface{}{1}, sql.Unique, 1, 2, 3)
  _ = mgr.SetAdapter(1)
  _, _ = mgr.ToSql()
}
