package nodes

import (
  "github.com/chuckpreslar/codex/nodes"
  "testing"
)

func TestOr(t *testing.T) {
  or := nodes.Or(1, 2)

  // The following struct members should exist.
  _ = or.Left
  _ = or.Right

  // The following receiver methods should exist.
  _ = or.Or(1)
  _ = or.And(1)
  _ = or.Not()
}
