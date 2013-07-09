package visitors

import (
  "github.com/chuckpreslar/codex/tree/nodes"
  "github.com/chuckpreslar/codex/tree/visitors"
  "testing"
)

var postgres = &visitors.PostgresVisitor{&visitors.ToSqlVisitor{}}

func TestPostgresLike(t *testing.T) {
  like := nodes.Like(1, 2)
  expected := "1 ILIKE 2"
  if got, _ := postgres.Accept(like); expected != got {
    t.Errorf("TestLike was expected to return %s, got %s", expected, got)
  }
}

func TestPostgresUnike(t *testing.T) {
  unlike := nodes.Unlike(1, 2)
  expected := "1 NOT ILIKE 2"
  if got, _ := postgres.Accept(unlike); expected != got {
    t.Errorf("TestUnlike was expected to return %s, got %s", expected, got)
  }
}
