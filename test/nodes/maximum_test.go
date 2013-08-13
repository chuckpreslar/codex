package nodes

import (
  "github.com/chuckpreslar/codex/nodes"
  "testing"
)

func TestMaximum(t *testing.T) {
  max := nodes.Maximum(1, 2, 3, 4)

  // The following struct members should exist.
  _ = max.Expressions
  _ = max.Alias
  _ = max.Distinct

  // The following receiver methods should exist.
  _ = max.And(1)
  _ = max.Or(1)
  _ = max.Eq(1)
  _ = max.Neq(1)
  _ = max.Gt(1)
  _ = max.Gte(1)
  _ = max.Lt(1)
  _ = max.Lte(1)
  _ = max.Like(1)
  _ = max.Unlike(1)
}
