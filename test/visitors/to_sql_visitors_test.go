package visitors

import (
  "github.com/chuckpreslar/codex/tree/visitors"
  "testing"
)

var sql = new(visitors.ToSqlVisitor)

func TestVisitString(t *testing.T) {
  value, expected := `test`, `'test'`
  if got := sql.Accept(value); expected != got {
    t.Errorf("VisitString was expected to return %s, got %s", expected, got)
  }
}

func TestVisitInteger(t *testing.T) {
  value, expected := 0, `0`
  if got := sql.Accept(value); expected != got {
    t.Errorf("VisitInteger was expected to return %s, got %s", expected, got)
  }
}

func TestVisitFloat(t *testing.T) {
  value, expected := 0.25, `0.25`
  if got := sql.Accept(value); expected != got {
    t.Errorf("VisitFloat was expected to return %s, got %s", expected, got)
  }
}

func TestVisitBool(t *testing.T) {
  value, expected := true, `'true'`
  if got := sql.Accept(value); expected != got {
    t.Errorf("VisitBool was expected to return %s, got %s", expected, got)
  }
}
