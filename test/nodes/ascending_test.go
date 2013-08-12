package nodes

import (
  "github.com/chuckpreslar/codex/nodes"
  "testing"
)

func TestAscending(t *testing.T) {
  asc := nodes.Ascending(1)

  // The following struct members should exist.
  _ = asc.Expr

  // The following receiver methods should exist.
  _ = asc.Or(1)
  _ = asc.And(1)
  _ = asc.Not()
}
