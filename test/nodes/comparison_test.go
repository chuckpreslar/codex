package nodes

import (
  "github.com/chuckpreslar/codex/tree/nodes"
  "testing"
)

func TestEqual(t *testing.T) {
  eq := nodes.Equal(1, 2)

  // The following struct members should exist.
  _ = eq.Left
  _ = eq.Right

  // The following receiver methods should exist.
  _ = eq.Or(1)
  _ = eq.And(1)
  _ = eq.Not()
}

func TestNotEqual(t *testing.T) {
  neq := nodes.NotEqual(1, 2)

  // The following struct members should exist.
  _ = neq.Left
  _ = neq.Right

  // The following receiver methods should exist.
  _ = neq.Or(1)
  _ = neq.And(1)
  _ = neq.Not()
}

func TestGreaterThan(t *testing.T) {
  gt := nodes.GreaterThan(1, 2)

  // The following struct members should exist.
  _ = gt.Left
  _ = gt.Right

  // The following receiver methods should exist.
  _ = gt.Or(1)
  _ = gt.And(1)
  _ = gt.Not()
}

func TestGreaterThanOrEqual(t *testing.T) {
  gte := nodes.GreaterThanOrEqual(1, 2)

  // The following struct members should exist.
  _ = gte.Left
  _ = gte.Right

  // The following receiver methods should exist.
  _ = gte.Or(1)
  _ = gte.And(1)
  _ = gte.Not()
}

func TestLessThan(t *testing.T) {
  lt := nodes.LessThan(1, 2)

  // The following struct members should exist.
  _ = lt.Left
  _ = lt.Right

  // The following receiver methods should exist.
  _ = lt.Or(1)
  _ = lt.And(1)
  _ = lt.Not()
}

func TestLessThanOrEqual(t *testing.T) {
  lte := nodes.LessThanOrEqual(1, 2)

  // The following struct members should exist.
  _ = lte.Left
  _ = lte.Right

  // The following receiver methods should exist.
  _ = lte.Or(1)
  _ = lte.And(1)
  _ = lte.Not()
}

func TestLike(t *testing.T) {
  like := nodes.Like(1, 2)

  // The following struct members should exist.
  _ = like.Left
  _ = like.Right

  // The following receiver methods should exist.
  _ = like.Or(1)
  _ = like.And(1)
  _ = like.Not()
}

func TestUnlike(t *testing.T) {
  unlike := nodes.Unlike(1, 2)

  // The following struct members should exist.
  _ = unlike.Left
  _ = unlike.Right

  // The following receiver methods should exist.
  _ = unlike.Or(1)
  _ = unlike.And(1)
  _ = unlike.Not()
}
