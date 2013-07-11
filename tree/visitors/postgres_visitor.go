package visitors

import (
  "errors"
  "fmt"
  "github.com/chuckpreslar/codex/tree/nodes"
)

type PostgresVisitor struct {
  *ToSqlVisitor
}

func (self *PostgresVisitor) Accept(o interface{}) (result string, err error) {
  defer func() {
    if r := recover(); r != nil {
      err = errors.New(fmt.Sprintf("%v", r))
    }
  }()

  result = self.Visit(o, self)

  return
}

func (self *PostgresVisitor) VisitLike(o *nodes.LikeNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v ILIKE %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *PostgresVisitor) VisitUnlike(o *nodes.UnlikeNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v NOT ILIKE %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}
