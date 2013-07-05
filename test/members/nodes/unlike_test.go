package nodes

import (
  "librarian/tree/members/nodes"
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
  if unlike != or.Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", unlike, or.Left)
  } else if other != or.Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, or.Right)
  }
}

func TestUnlikeAnd(t *testing.T) {
  left, right := 1, 2
  unlike := &nodes.Unlike{left, right}
  other := 3
  and := unlike.And(other)
  if unlike != and.Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", unlike, and.Left)
  } else if other != and.Right {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", other, and.Right)
  }
}
