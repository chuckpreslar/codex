package nodes

import (
  "librarian/tree/members/nodes"
  "testing"
)

func TestNotEqual(t *testing.T) {
  left, right := 1, 2
  equal := &nodes.NotEqual{&nodes.Binary{left, right}}
  if left != equal.Left {
    t.Errorf("Expect Left NotEqual leaf to equal %v, got %v.", left, equal.Left)
  } else if right != equal.Right {
    t.Errorf("Expect Right NotEqual leaf to equal %v, got %v.", right, equal.Right)
  }
}
