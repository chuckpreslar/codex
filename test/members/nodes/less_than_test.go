package nodes

import (
  "librarian/tree/members/nodes"
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
  if lessThan != or.Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", lessThan, or.Left)
  } else if other != or.Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, or.Right)
  }
}

func TestLessThanAnd(t *testing.T) {
  left, right := 1, 2
  lessThan := &nodes.LessThan{left, right}
  other := 3
  and := lessThan.And(other)
  if lessThan != and.Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", lessThan, and.Left)
  } else if other != and.Right {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", other, and.Right)
  }
}
