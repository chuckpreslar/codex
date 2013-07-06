package nodes

import (
  "librarian/tree/nodes"
  "testing"
)

func TestUnlike(t *testing.T) {
  left, right := 1, 2
  unlike := &nodes.Unlike{left, right}
  if left != unlike.Left {
    t.Errorf("Expect Left Unlike leaf to equal %v, got %v.", left, unlike.Left)
  } else if right != unlike.Right {
    t.Errorf("Expect Right Unlike leaf to equal %v, got %v.", right, unlike.Right)
  }
}

func TestUnlikeOr(t *testing.T) {
  left, right := 1, 2
  unlike := &nodes.Unlike{left, right}
  other := 3
  or := unlike.Or(other)
  if unlike != or.Expr.(*nodes.Or).Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", unlike, or.Expr.(*nodes.Or).Left)
  } else if other != or.Expr.(*nodes.Or).Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, or.Expr.(*nodes.Or).Right)
  }
}

func TestUnlikeAnd(t *testing.T) {
  left, right := 1, 2
  unlike := &nodes.Unlike{left, right}
  other := 3
  and := unlike.And(other)
  if unlike != and.Expr.(*nodes.And).Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", unlike, and.Expr.(*nodes.And).Left)
  } else if other != and.Expr.(*nodes.And).Right {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", other, and.Expr.(*nodes.And).Right)
  }
}
