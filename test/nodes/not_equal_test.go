package nodes

import (
  "codex/tree/nodes"
  "testing"
)

func TestNotEqual(t *testing.T) {
  left, right := 1, 2
  notEqual := &nodes.NotEqual{left, right}
  if left != notEqual.Left {
    t.Errorf("Expect Left NotEqual leaf to equal %v, got %v.", left, notEqual.Left)
  } else if right != notEqual.Right {
    t.Errorf("Expect Right NotEqual leaf to equal %v, got %v.", right, notEqual.Right)
  }
}

func TestNotEqualOr(t *testing.T) {
  left, right := 1, 2
  notEqual := &nodes.NotEqual{left, right}
  other := 3
  or := notEqual.Or(other)
  if notEqual != or.Expr.(*nodes.Or).Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", notEqual, notEqual.Left)
  } else if other != or.Expr.(*nodes.Or).Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, notEqual.Right)
  }
}

func TestNotEqualAnd(t *testing.T) {
  left, right := 1, 2
  notEqual := &nodes.NotEqual{left, right}
  other := 3
  and := notEqual.And(other)
  if notEqual != and.Expr.(*nodes.And).Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", notEqual, notEqual.Left)
  } else if other != and.Expr.(*nodes.And).Right {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", other, notEqual.Right)
  }
}
