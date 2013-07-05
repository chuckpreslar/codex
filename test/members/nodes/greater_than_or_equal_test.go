package nodes

import (
  "librarian/tree/members/nodes"
  "testing"
)

func TestGreaterThanOrEqual(t *testing.T) {
  left, right := 1, 2
  greaterThanOrEqual := &nodes.GreaterThanOrEqual{left, right}
  if left != greaterThanOrEqual.Left {
    t.Errorf("Expect Left GreaterThanOrEqual leaf to equal %v, got %v.", left, greaterThanOrEqual.Left)
  } else if right != greaterThanOrEqual.Right {
    t.Errorf("Expect Right GreaterThanOrEqual leaf to equal %v, got %v.", right, greaterThanOrEqual.Right)
  }
}

func TestGreaterThanOrEqualOr(t *testing.T) {
  left, right := 1, 2
  greaterThanOrEqual := &nodes.GreaterThanOrEqual{left, right}
  other := 3
  or := greaterThanOrEqual.Or(other)
  if greaterThanOrEqual != or.Expr.(*nodes.Or).Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", greaterThanOrEqual, or.Expr.(*nodes.Or).Left)
  } else if other != or.Expr.(*nodes.Or).Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, or.Expr.(*nodes.Or).Right)
  }
}

func TestGreaterThanOrEqualAnd(t *testing.T) {
  left, right := 1, 2
  greaterThanOrEqual := &nodes.GreaterThanOrEqual{left, right}
  other := 3
  and := greaterThanOrEqual.And(other)
  if greaterThanOrEqual != and.Expr.(*nodes.And).Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", greaterThanOrEqual, and.Expr.(*nodes.And).Left)
  } else if other != and.Expr.(*nodes.And).Right  {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", other, and.Expr.(*nodes.And).Right )
  }
}
