package librarian

import (
  "testing"
)

var users Table = NewTable("users")

func TestSelect(t *testing.T) {
  stmt := Search(users).Select("id", users("email"))
  if 2 != len(stmt.projections) {
    t.Errorf("Expected 2 projections to be added to the Select Statement, found %d.\n", len(stmt.projections))
  }
}

func TestWhere(t *testing.T) {
  stmt := Search(users).Where(users("id").As("user_id").Eq("1"))
  if 1 != len(stmt.filters) {
    t.Errorf("Expected 1 filter to be added to the Select Statement, found %d.\n", len(stmt.filters))
  }
}
