package nodes

import (
  "github.com/chuckpreslar/codex/tree/nodes"
  "testing"
)

func TestAttribute(t *testing.T) {
  attr := nodes.Attribute("column", nodes.Relation("table"))

  // The following struct members should exist.
  _ = attr.Name
  _ = attr.Relation

  // The following receiver methods should exist.
  _ = attr.And(1)
  _ = attr.Or(1)
  _ = attr.Eq(1)
  _ = attr.Neq(1)
  _ = attr.Gt(1)
  _ = attr.Gte(1)
  _ = attr.Lt(1)
  _ = attr.Lte(1)
  _ = attr.Like(1)
  _ = attr.Unlike(1)
  _ = attr.Asc()
  _ = attr.Desc()
}
