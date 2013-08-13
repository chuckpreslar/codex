package visitors

import (
  "github.com/chuckpreslar/codex/nodes"
  "github.com/chuckpreslar/codex/visitors"
  "testing"
)

var postgres = &visitors.PostgresVisitor{&visitors.ToSqlVisitor{}, 0}

func TestPostgresBinding(t *testing.T) {
  binding := nodes.Binding()
  expected := "$1"
  if got, _ := postgres.Accept(binding); expected != got {
    t.Errorf("TestPostgresBinding was expected to return %s, got %s", expected, got)
  }
}

func TestPostgresLike(t *testing.T) {
  like := nodes.Like(1, 2)
  expected := "1 ILIKE 2"
  if got, _ := postgres.Accept(like); expected != got {
    t.Errorf("TestPostgresLike was expected to return %s, got %s", expected, got)
  }
}

func TestPostgresUnike(t *testing.T) {
  unlike := nodes.Unlike(1, 2)
  expected := "1 NOT ILIKE 2"
  if got, _ := postgres.Accept(unlike); expected != got {
    t.Errorf("TestPostgresUnike was expected to return %s, got %s", expected, got)
  }
}
