package nodes

import (
  "codex/tree/nodes"
  "testing"
)

func TestLessThan(t *testing.T) {
  left, right := 1, 2
  lessThan := &nodes.LessThan{left, right}
  if left != lessThan.Left {
    t.Errorf("Expect Left LessThan leaf to equal %v, got %v.", left, lessThan.Left)
  } else if right != lessThan.Right {
    t.Errorf("Expect Right LessThan leaf to equal %v, got %v.", right, lessThan.Right)
  }
}

func TestLessThanOr(t *testing.T) {
  left, right := 1, 2
  lessThan := &nodes.LessThan{left, right}
  other := 3
  or := lessThan.Or(other)
  if lessThan != or.Expr.(*nodes.Or).Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", lessThan, or.Expr.(*nodes.Or).Left)
  } else if other != or.Expr.(*nodes.Or).Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, or.Expr.(*nodes.Or).Right)
  }
}

func TestLessThanAnd(t *testing.T) {
  left, right := 1, 2
  lessThan := &nodes.LessThan{left, right}
  other := 3
  and := lessThan.And(other)
  if lessThan != and.Expr.(*nodes.And).Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", lessThan, and.Expr.(*nodes.And).Left)
  } else if other != and.Expr.(*nodes.And).Right {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", other, and.Expr.(*nodes.And).Right)
  }
}
