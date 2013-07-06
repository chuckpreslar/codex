package nodes

import (
  "codex/tree/nodes"
  "testing"
)

func TestLike(t *testing.T) {
  left, right := 1, 2
  like := &nodes.Like{left, right}
  if left != like.Left {
    t.Errorf("Expect Left Like leaf to equal %v, got %v.", left, like.Left)
  } else if right != like.Right {
    t.Errorf("Expect Right Like leaf to equal %v, got %v.", right, like.Right)
  }
}

func TestLikeOr(t *testing.T) {
  left, right := 1, 2
  like := &nodes.Like{left, right}
  other := 3
  or := like.Or(other)
  if like != or.Expr.(*nodes.Or).Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", like, or.Expr.(*nodes.Or).Left)
  } else if other != or.Expr.(*nodes.Or).Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, or.Expr.(*nodes.Or).Right)
  }
}

func TestLikeAnd(t *testing.T) {
  left, right := 1, 2
  like := &nodes.Like{left, right}
  other := 3
  and := like.And(other)
  if like != and.Expr.(*nodes.And).Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", like, and.Expr.(*nodes.And).Left)
  } else if other != and.Expr.(*nodes.And).Right {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", other, and.Expr.(*nodes.And).Right)
  }
}
