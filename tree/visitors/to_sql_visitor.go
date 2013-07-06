package visitors

import (
  "fmt"
  "librarian/tree/nodes"
  "strings"
)

const (
  SPACE = ` `
  COMMA = `, `
)

type ToSqlVisitor struct{}

func (visitor *ToSqlVisitor) Accept(o interface{}) string {
  return visitor.Visit(o)
}

func (visitor *ToSqlVisitor) Visit(o interface{}) string {
  switch o.(type) {
  // Comparison visitors.
  case *nodes.Equal:
    return visitor.VisitEqual(o.(*nodes.Equal))
  case *nodes.NotEqual:
    return visitor.VisitNotEqual(o.(*nodes.NotEqual))
  case *nodes.GreaterThan:
    return visitor.VisitGreaterThan(o.(*nodes.GreaterThan))
  case *nodes.GreaterThanOrEqual:
    return visitor.VisitGreaterThanOrEqual(o.(*nodes.GreaterThanOrEqual))
  case *nodes.LessThan:
    return visitor.VisitLessThan(o.(*nodes.LessThan))
  case *nodes.LessThanOrEqual:
    return visitor.VisitLessThanOrEqual(o.(*nodes.LessThanOrEqual))
  case *nodes.Like:
    return visitor.VisitLike(o.(*nodes.Like))
  case *nodes.Unlike:
    return visitor.VisitUnlike(o.(*nodes.Unlike))
  case *nodes.Or:
    return visitor.VisitOr(o.(*nodes.Or))
  case *nodes.And:
    return visitor.VisitAnd(o.(*nodes.And))
  // Begin SQL node visitors.
  case *nodes.Grouping:
    return visitor.VisitGrouping(o.(*nodes.Grouping))
  case *nodes.Not:
    return visitor.VisitNot(o.(*nodes.Not))
  case *nodes.InnerJoin:
    return visitor.VisitInnerJoin(o.(*nodes.InnerJoin))
  case *nodes.OuterJoin:
    return visitor.VisitOuterJoin(o.(*nodes.OuterJoin))
  case *nodes.On:
    return visitor.VisitOn(o.(*nodes.On))
  // Standard type visitors.
  case string:
    return visitor.VisitString(o)
  case int, int16, int32, int64:
    return visitor.VisitInteger(o)
  case float32, float64:
    return visitor.VisitFloat(o)
  case bool:
    return visitor.VisitBool(o)
  default:
    panic(fmt.Sprintf("No visitor method for <%T>.", o))
  }
}

// Being comparison node visitors.

func (visitor *ToSqlVisitor) VisitEqual(o *nodes.Equal) string {
  if nil == o.Right {
    return fmt.Sprintf("%v IS NULL", visitor.Visit(o.Left))
  }

  return fmt.Sprintf("%v = %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitNotEqual(o *nodes.NotEqual) string {
  if nil == o.Right {
    return fmt.Sprintf("%v IS NOT NULL", visitor.Visit(o.Left))
  }

  return fmt.Sprintf("%v != %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitGreaterThan(o *nodes.GreaterThan) string {
  return fmt.Sprintf("%v > %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitGreaterThanOrEqual(o *nodes.GreaterThanOrEqual) string {
  return fmt.Sprintf("%v >= %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitLessThan(o *nodes.LessThan) string {
  return fmt.Sprintf("%v < %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitLessThanOrEqual(o *nodes.LessThanOrEqual) string {
  return fmt.Sprintf("%v <= %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitLike(o *nodes.Like) string {
  return fmt.Sprintf("%v LIKE %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitUnlike(o *nodes.Unlike) string {
  return fmt.Sprintf("%v NOT LIKE %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitOr(o *nodes.Or) string {
  return fmt.Sprintf("%v OR %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitAnd(o *nodes.And) string {
  return fmt.Sprintf("%v AND %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

// End comparison node visitors.

// Begin SQL node visitors.

func (visitor *ToSqlVisitor) VisitGrouping(o *nodes.Grouping) string {
  return fmt.Sprintf("(%v)", visitor.Visit(o.Expr))
}

func (visitor *ToSqlVisitor) VisitNot(o *nodes.Not) string {
  return fmt.Sprintf("NOT (%v)", visitor.Visit(o.Expr))
}

func (visitor *ToSqlVisitor) VisitInnerJoin(o *nodes.InnerJoin) string {
  str := fmt.Sprintf("INNER JOIN %v", visitor.Visit(o.Left))
  if nil != o.Right {
    str = fmt.Sprintf("%v%v%v", str, SPACE, visitor.Visit(o.Right))
  }
  return str
}

func (visitor *ToSqlVisitor) VisitOuterJoin(o *nodes.OuterJoin) string {
  return fmt.Sprintf("LEFT OUTER JOIN %v %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitOn(o *nodes.On) string {
  return fmt.Sprintf("ON %v", visitor.Visit(o.Expr))
}

// End SQL node visitors.

// Begin standard type visitors.

func (visitor *ToSqlVisitor) VisitString(o interface{}) string {
  return fmt.Sprintf(`'%s'`, o)
}

func (visitor *ToSqlVisitor) VisitInteger(o interface{}) string {
  return fmt.Sprintf(`%d`, o)
}

func (visitor *ToSqlVisitor) VisitFloat(o interface{}) string {
  return strings.TrimRight(fmt.Sprintf(`%v`, o), `0`)
}

func (visitor *ToSqlVisitor) VisitBool(o interface{}) string {
  return fmt.Sprintf(`'%t'`, o)
}

// End standard Type visitors.
