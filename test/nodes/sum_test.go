package nodes

import (
  "github.com/chuckpreslar/codex/nodes"
  "testing"
)

func TestSum(t *testing.T) {
  sum := nodes.Sum(1, 2, 3, 4)

  // The following struct members should exist.
  _ = sum.Expressions
  _ = sum.Alias
  _ = sum.Distinct

  // The following receiver methods should exist.
  _ = sum.And(1)
  _ = sum.Or(1)
  _ = sum.Eq(1)
  _ = sum.Neq(1)
  _ = sum.Gt(1)
  _ = sum.Gte(1)
  _ = sum.Lt(1)
  _ = sum.Lte(1)
  _ = sum.Like(1)
  _ = sum.Unlike(1)
}
