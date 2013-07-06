package nodes

import (
  "codex/tree/nodes"
  "testing"
)

func TestBinary(t *testing.T) {
  left, right := 1, 2
  binary := &nodes.Binary{left, right}
  if left != binary.Left {
    t.Errorf("Expect Left Binary leaf to equal %v, got %v.", left, binary.Left)
  } else if right != binary.Right {
    t.Errorf("Expect Right Binary leaf to equal %v, got %v.", right, binary.Right)
  }
}
