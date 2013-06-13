package librarian

import (
  "testing"
)

func TestSelect(t *testing.T) {
  users := NewTable("users")
  stmt := Search(users).Select("id", users("email"))
  if 2 != len(stmt.projections) {
    t.Errorf("Expected 2 projections to be added to the Select Statement, found %d.\n", len(stmt.projections))
  }
}
