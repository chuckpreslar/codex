package visitors

import (
  "fmt"
  "librarian/tree/nodes"
  "strings"
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
    panic(fmt.Sprintf("No visitor method for <%T>.", o, o))
  }
}

// Being comparison visitors.

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

// End comparison visitors.

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
