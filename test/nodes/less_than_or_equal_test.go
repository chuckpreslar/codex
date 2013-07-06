package nodes

import (
  "librarian/tree/nodes"
  "testing"
)

func TestLessThanOrEqual(t *testing.T) {
  left, right := 1, 2
  lessThanOrEqual := &nodes.LessThanOrEqual{left, right}
  if left != lessThanOrEqual.Left {
    t.Errorf("Expect Left LessThanOrEqual leaf to equal %v, got %v.", left, lessThanOrEqual.Left)
  } else if right != lessThanOrEqual.Right {
    t.Errorf("Expect Right LessThanOrEqual leaf to equal %v, got %v.", right, lessThanOrEqual.Right)
  }
}

func TestLessThanOrEqualOr(t *testing.T) {
  left, right := 1, 2
  lessThanOrEqual := &nodes.LessThanOrEqual{left, right}
  other := 3
  or := lessThanOrEqual.Or(other)
  if lessThanOrEqual != or.Expr.(*nodes.Or).Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", lessThanOrEqual, or.Expr.(*nodes.Or).Left)
  } else if other != or.Expr.(*nodes.Or).Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, or.Expr.(*nodes.Or).Right)
  }
}

func TestLessThanOrEqualAnd(t *testing.T) {
  left, right := 1, 2
  lessThanOrEqual := &nodes.LessThanOrEqual{left, right}
  other := 3
  and := lessThanOrEqual.And(other)
  if lessThanOrEqual != and.Expr.(*nodes.And).Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", lessThanOrEqual, and.Expr.(*nodes.And).Left)
  } else if other != and.Expr.(*nodes.And).Right {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", other, and.Expr.(*nodes.And).Right)
  }
}
