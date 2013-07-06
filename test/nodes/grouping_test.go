package nodes

import (
  "librarian/tree/nodes"
  "testing"
)

func TestGrouping(t *testing.T) {
  expected := 1
  grouping := &nodes.Grouping{expected}
  if expected != grouping.Expr {
    t.Errorf("Expected Expr of Grouping to be %v, got %v.", expected, grouping.Expr)
  }
}
