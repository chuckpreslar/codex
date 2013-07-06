package nodes

import (
  "librarian/tree/nodes"
  "testing"
)

func TestOr(t *testing.T) {
  left, right := 1, 2
  or := &nodes.Or{left, right}
  if left != or.Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", left, or.Left)
  } else if right != or.Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", right, or.Right)
  }
}
