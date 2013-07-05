package nodes

import (
  "librarian/tree/members/nodes"
  "testing"
)

func TestEqual(t *testing.T) {
  left, right := 1, 2
  equal := &nodes.Equal{left, right}
  if left != equal.Left {
    t.Errorf("Expect Left Equal leaf to equal %v, got %v.", left, equal.Left)
  } else if right != equal.Right {
    t.Errorf("Expect Right Equal leaf to equal %v, got %v.", right, equal.Right)
  }
}

func TestEqualOr(t *testing.T) {
  left, right := 1, 2
  equal := &nodes.Equal{left, right}
  other := 3
  or := equal.Or(other)
  if equal != or.Expr.(*nodes.Or).Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", equal, equal.Left)
  } else if other != or.Expr.(*nodes.Or).Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, equal.Right)
  }
}

func TestEqualAnd(t *testing.T) {
  left, right := 1, 2
  equal := &nodes.Equal{left, right}
  other := 3
  and := equal.And(other)
  if equal != and.Expr.(*nodes.And).Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", equal, equal.Left)
  } else if other != and.Expr.(*nodes.And).Right {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", other, equal.Right)
  }
}
