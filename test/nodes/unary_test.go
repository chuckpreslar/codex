package nodes

import (
  "librarian/tree/nodes"
  "testing"
)

func TestUnary(t *testing.T) {
  expr := 1
  unary := &nodes.Unary{expr}
  if expr != unary.Expr {
    t.Errorf("Expect Unary Expr leaf to equal %v, got %v.", expr, unary.Expr)
  }
}
