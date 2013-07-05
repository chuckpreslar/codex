package nodes

import (
  "librarian/tree/members/nodes"
  "testing"
)

func TestLessThanOrEqual(t *testing.T) {
  left, right := 1, 2
  lessThanOrEqual := &nodes.LessThanOrEqual{&nodes.Binary{left, right}}
  if left != lessThanOrEqual.Left {
    t.Errorf("Expect Left LessThanOrEqual leaf to equal %v, got %v.", left, lessThanOrEqual.Left)
  } else if right != lessThanOrEqual.Right {
    t.Errorf("Expect Right LessThanOrEqual leaf to equal %v, got %v.", right, lessThanOrEqual.Right)
  }
}

func TestLessThanOrEqualOr(t *testing.T) {
  left, right := 1, 2
  lessThanOrEqual := &nodes.LessThanOrEqual{&nodes.Binary{left, right}}
  other := 3
  or := lessThanOrEqual.Or(other)
  if lessThanOrEqual != or.Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", lessThanOrEqual, or.Left)
  } else if other != or.Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, or.Right)
  }
}

func TestLessThanOrEqualAnd(t *testing.T) {
  left, right := 1, 2
  lessThanOrEqual := &nodes.LessThanOrEqual{&nodes.Binary{left, right}}
  other := 3
  and := lessThanOrEqual.And(other)
  if lessThanOrEqual != and.Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", lessThanOrEqual, and.Left)
  } else if other != and.Right {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", other, and.Right)
  }
}
