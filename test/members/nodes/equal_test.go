package nodes

import (
  "librarian/tree/members/nodes"
  "testing"
)

func TestEqual(t *testing.T) {
  left, right := 1, 2
  equal := &nodes.Equal{&nodes.Binary{left, right}}
  if left != equal.Left {
    t.Errorf("Expect Left Equal leaf to equal %v, got %v.", left, equal.Left)
  } else if right != equal.Right {
    t.Errorf("Expect Right Equal leaf to equal %v, got %v.", right, equal.Right)
  }
}
