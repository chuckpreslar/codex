package nodes

import (
  "librarian/tree/members/nodes"
  "testing"
)

func TestGreaterThan(t *testing.T) {
  left, right := 1, 2
  greaterThan := &nodes.GreaterThan{&nodes.Binary{left, right}}
  if left != greaterThan.Left {
    t.Errorf("Expect Left GreaterThan leaf to equal %v, got %v.", left, greaterThan.Left)
  } else if right != greaterThan.Right {
    t.Errorf("Expect Right GreaterThan leaf to equal %v, got %v.", right, greaterThan.Right)
  }
}

func TestGreaterThanOr(t *testing.T) {
  left, right := 1, 2
  greaterThan := &nodes.GreaterThan{&nodes.Binary{left, right}}
  other := 3
  or := greaterThan.Or(other)
  if greaterThan != or.Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", greaterThan, or.Left)
  } else if other != or.Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, or.Right)
  }
}

func TestGreaterThanAnd(t *testing.T) {
  left, right := 1, 2
  greaterThan := &nodes.GreaterThan{&nodes.Binary{left, right}}
  other := 3
  and := greaterThan.And(other)
  if greaterThan != and.Left {
    t.Errorf("Expect Left And leaf to equal %v, got %v.", greaterThan, and.Left)
  } else if other != and.Right {
    t.Errorf("Expect Right And leaf to equal %v, got %v.", other, and.Right)
  }
}
