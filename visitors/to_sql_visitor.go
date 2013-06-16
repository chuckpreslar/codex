package visitors

import (
  "fmt"
  "librarian/nodes"
)

type ToSqlVisitor struct {
  ast nodes.NodeInterface
}

func (visitor *ToSqlVisitor) Accept(ast nodes.NodeInterface) string {
  visitor.ast = ast
  return visitor.visit(ast)
}

func (visitor *ToSqlVisitor) visit(item interface{}) string {
  switch item.(type) {
  case *nodes.EqNode:
    return visitor.visitEqNode(item.(*nodes.EqNode))
  case *nodes.AttributeNode:
    return visitor.visitAttributeNode(item.(*nodes.AttributeNode))
  case string:
    return visitor.visitString(item.(string))
  }
  panic("Unimplemented.")
}

func (visitor *ToSqlVisitor) visitEqNode(eq *nodes.EqNode) string {
  return fmt.Sprintf("%v = %v", visitor.visit(eq.Left()), tag(visitor.visit(eq.Right())))
}

func (visitor *ToSqlVisitor) visitAttributeNode(attribute *nodes.AttributeNode) string {
  return fmt.Sprintf("%s.%s", quote(visitor.visit(attribute.Left())), quote(visitor.visit(attribute.Right())))
}

func (visitor *ToSqlVisitor) visitString(str string) string {
  return str
}

func ToSql() *ToSqlVisitor {
  return &ToSqlVisitor{}
}

func quote(value interface{}) string {
  return fmt.Sprintf(`"%v"`, value)
}

func tag(value interface{}) string {
  switch value.(type) {
  case string:
    return fmt.Sprintf(`'%v'`, value)
  case bool:
    return fmt.Sprintf(`'%v'`, value)
  default:
    return fmt.Sprintf(`%v`, value)
  }
}
