package visitors

import (
  "github.com/chuckpreslar/codex/tree/nodes"
  "fmt"
)

type PostgresVisitor struct {
  *ToSqlVisitor
}

func (postgres *PostgresVisitor) Accept(o interface{}) string {
  return postgres.Visit(o, postgres)
}

func (postgres *PostgresVisitor) VisitLike(o *nodes.Like, visitor VisitorInterface) string {
  return fmt.Sprintf("%v ILIKE %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (postgres *PostgresVisitor) VisitUnlike(o *nodes.Unlike, visitor VisitorInterface) string {
  return fmt.Sprintf("%v NOT ILIKE %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}
