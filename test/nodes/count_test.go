package nodes

import (
  "github.com/chuckpreslar/codex/tree/nodes"
  "testing"
)

func TestCount(t *testing.T) {
  count := nodes.Count(1, 2, 3, 4)

  // The following struct members should exist.
  _ = count.Expressions
  _ = count.Alias
  _ = count.Distinct

  // The following receiver methods should exist.
  _ = count.And(1)
  _ = count.Or(1)
  _ = count.Eq(1)
  _ = count.Neq(1)
  _ = count.Gt(1)
  _ = count.Gte(1)
  _ = count.Lt(1)
  _ = count.Lte(1)
  _ = count.Like(1)
  _ = count.Unlike(1)
}
