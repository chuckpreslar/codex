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
  case *nodes.NeqNode:
    return visitor.visitNeqNode(item.(*nodes.NeqNode))
  case *nodes.GtNode:
    return visitor.visitGtNode(item.(*nodes.GtNode))
  case *nodes.GteNode:
    return visitor.visitGteNode(item.(*nodes.GteNode))
  case *nodes.LtNode:
    return visitor.visitLtNode(item.(*nodes.LtNode))
  case *nodes.LteNode:
    return visitor.visitLteNode(item.(*nodes.LteNode))
  case *nodes.AttributeNode:
    return visitor.visitAttributeNode(item.(*nodes.AttributeNode))
  case string:
    return visitor.visitString(item.(string))
  case int:
    return visitor.visitInt(item.(int))
  }
  panic("Unimplemented.")
}

func (visitor *ToSqlVisitor) visitEqNode(eq *nodes.EqNode) string {
  if nil == eq.Right() {
    return fmt.Sprintf("%v IS NULL", visitor.visit(eq.Left()))
  }
  return fmt.Sprintf("%v = %v", visitor.visit(eq.Left()), tag(visitor.visit(eq.Right())))
}

func (visitor *ToSqlVisitor) visitNeqNode(neq *nodes.NeqNode) string {
  if nil == neq.Right() {
    return fmt.Sprintf("%v IS NOT NULL", visitor.visit(neq.Left()))
  }
  return fmt.Sprintf("%v != %v", visitor.visit(neq.Left()), tag(visitor.visit(neq.Right())))
}

func (visitor *ToSqlVisitor) visitGtNode(gt *nodes.GtNode) string {
  return fmt.Sprintf("%v > %v", visitor.visit(gt.Left()), visitor.visit(gt.Right()))
}

func (visitor *ToSqlVisitor) visitGteNode(gte *nodes.GteNode) string {
  return fmt.Sprintf("%v >= %v", visitor.visit(gte.Left()), visitor.visit(gte.Right()))
}

func (visitor *ToSqlVisitor) visitLtNode(lt *nodes.LtNode) string {
  return fmt.Sprintf("%v < %v", visitor.visit(lt.Left()), visitor.visit(lt.Right()))
}

func (visitor *ToSqlVisitor) visitLteNode(lte *nodes.LteNode) string {
  return fmt.Sprintf("%v <= %v", visitor.visit(lte.Left()), visitor.visit(lte.Right()))
}

func (visitor *ToSqlVisitor) visitAttributeNode(attribute *nodes.AttributeNode) string {
  return fmt.Sprintf("%s.%s", quote(visitor.visit(attribute.Left())), quote(visitor.visit(attribute.Right())))
}

func (visitor *ToSqlVisitor) visitString(str string) string {
  return str
}

func (visitor *ToSqlVisitor) visitInt(i int) string {
  return fmt.Sprintf("%d", i)
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
