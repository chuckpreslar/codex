// Package visitors provides AST visitors for the codex package.
package visitors

import (
  "errors"
  "fmt"
  "github.com/chuckpreslar/codex/tree/nodes"
)

type PostgresVisitor struct {
  *ToSqlVisitor
  Bindings int // Number of bindings used in parameterization
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

func (self *PostgresVisitor) VisitBinding(o *nodes.BindingNode, visitor VisitorInterface) string {
  self.Bindings += 1
  return fmt.Sprintf("$%d", self.Bindings)
}

func (self *PostgresVisitor) VisitLike(o *nodes.LikeNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v ILIKE %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *PostgresVisitor) VisitUnlike(o *nodes.UnlikeNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v NOT ILIKE %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}
