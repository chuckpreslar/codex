package visitors

import (
  "librarian/tree/visitors"
  "testing"
)

func TestVisitString(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  value, expected := `test`, `'test'`
  if got := visitor.Accept(value); expected != got {
    t.Errorf("VisitString was expected to return %s, got %f", expected, got)
  }
}

func TestVisitInteger(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  value, expected := 0, `0`
  if got := visitor.Accept(value); expected != got {
    t.Errorf("VisitInteger was expected to return %s, got %f", expected, got)
  }
}

func TestVisitFloat(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  value, expected := 0.25, `0.25`
  if got := visitor.Accept(value); expected != got {
    t.Errorf("VisitFloat was expected to return %s, got %f", expected, got)
  }
}

func TestVisitBool(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  value, expected := true, `'true'`
  if got := visitor.Accept(value); expected != got {
    t.Errorf("VisitBool was expected to return %s, got %f", expected, got)
  }
}
