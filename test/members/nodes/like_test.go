package nodes

import (
  "librarian/tree/members/nodes"
  "testing"
)

func TestLike(t *testing.T) {
  left, right := 1, 2
  like := &nodes.Like{&nodes.Binary{left, right}}
  if left != like.Left {
    t.Errorf("Expect Left Like leaf to equal %v, got %v.", left, like.Left)
  } else if right != like.Right {
    t.Errorf("Expect Right Like leaf to equal %v, got %v.", right, like.Right)
  }
}

func TestLikeOr(t *testing.T) {
  left, right := 1, 2
  like := &nodes.Like{&nodes.Binary{left, right}}
  other := 3
  or := like.Or(other)
  if like != or.Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", like, or.Left)
  } else if other != or.Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, or.Right)
  }
}

func TestLikeAnd(t *testing.T) {
  left, right := 1, 2
  like := &nodes.Like{&nodes.Binary{left, right}}
  other := 3
  and := like.And(other)
  if like != and.Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", like, and.Left)
  } else if other != and.Right {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", other, and.Right)
  }
}
