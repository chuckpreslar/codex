package visitors

import (
  "github.com/chuckpreslar/codex/visitors"
  "testing"
)

var mysql = &visitors.MySqlVisitor{&visitors.ToSqlVisitor{}}

func TestQuoteColumnName(t *testing.T) {
  expected := "`column`"
  if got := mysql.QuoteColumnName("column"); expected != got {
    t.Errorf("TestQuoteColumnName was expected to return %s, got %s", expected, got)
  }
}

func TestQuoteTableName(t *testing.T) {
  expected := "`column`"
  if got := mysql.QuoteTableName("column"); expected != got {
    t.Errorf("TestQuoteColumnName was expected to return %s, got %s", expected, got)
  }
}
