package visitors

import (
  "fmt"
  "strings"
  _ "librarian/tree/nodes"
)

type ToSqlVisitor struct{}

func (visitor *ToSqlVisitor) Accept(o interface{}) string {
  return visitor.Visit(o)
}

func (visitor *ToSqlVisitor) Visit(o interface{}) string {
  switch o.(type) {
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
    panic(fmt.Sprintf("No visitor method for %v <%T>.", o, o))
  }
}

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
