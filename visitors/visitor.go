package visitors

import (
  "fmt"
  "librarian/nodes"
  "librarian/visitors/utils"
)

type Visitor interface {
  Accept(nodes.Node) string
  Visit(nodes.Node) string
  VisitSelectStatement(nodes.SelectStatement) string
  VisitProjections([]nodes.Projection) string
  VisitEqNode(nodes.EqNode) string
  VisitNeqNode(nodes.NeqNode) string
  VisitGtNode(nodes.GtNode) string
  VisitGteNode(nodes.GteNode) string
  VisitLtNode(nodes.LtNode) string
  VisitLteNode(nodes.LteNode) string
  VisitMatchesNode(nodes.MatchesNode) string
  VisitOrNode(nodes.OrNode) string
  VisitAttribute(nodes.Attribute) string
}

type BaseVisitor struct {
  ast nodes.Node
}

func (visitor BaseVisitor) Accept(node nodes.Node) string {
  visitor.ast = node
  return visitor.Visit(node)
}

func (visitor BaseVisitor) Visit(node nodes.Node) string {
  switch node.(type) {
  case *nodes.SelectStatement:
    return visitor.VisitSelectStatement((*node.(*nodes.SelectStatement)))
  case nodes.EqNode:
    return visitor.VisitEqNode(node.(nodes.EqNode))
  case nodes.NeqNode:
    return visitor.VisitNeqNode(node.(nodes.NeqNode))
  case nodes.GtNode:
    return visitor.VisitGtNode(node.(nodes.GtNode))
  case nodes.GteNode:
    return visitor.VisitGteNode(node.(nodes.GteNode))
  case nodes.LtNode:
    return visitor.VisitLtNode(node.(nodes.LtNode))
  case nodes.LteNode:
    return visitor.VisitLteNode(node.(nodes.LteNode))
  case nodes.MatchesNode:
    return visitor.VisitMatchesNode(node.(nodes.MatchesNode))
  case nodes.OrNode:
    return visitor.VisitOrNode(node.(nodes.OrNode))
  case nodes.Attribute:
    return visitor.VisitAttribute(node.(nodes.Attribute))
  }
  return ""
}

func (visitor BaseVisitor) VisitSelectStatement(node nodes.SelectStatement) string {
  sql := fmt.Sprintf("SELECT %s FROM %s",
    visitor.VisitProjections(node.Projections()), visitor.VisitReference(node.Reference()))
  fmt.Println(sql)
  return sql
}

func (visitor BaseVisitor) VisitProjections(projections []nodes.Projection) string {
  // if 0 == len(projections) {
  //   return visitor.quotedAttribute(visitor.ast.(nodes.SelectStatement).Reference().Name(), "*")
  // }
  columns := make([]interface{}, len(projections))
  for index, column := range projections {
    columns[index] = visitor.quotedAttribute(column.(nodes.Attribute).Reference().Name(),
      column.(nodes.Attribute).Name())
  }
  return utils.Join(columns, ", ")
}

func (visitor BaseVisitor) VisitReference(node *nodes.Reference) string {
  return visitor.quote(node.Name())
}

func (visitor BaseVisitor) VisitEqNode(node nodes.EqNode) string {
  return ""
}

func (visitor BaseVisitor) VisitNeqNode(node nodes.NeqNode) string {
  return ""
}

func (visitor BaseVisitor) VisitGtNode(node nodes.GtNode) string {
  return ""
}

func (visitor BaseVisitor) VisitGteNode(node nodes.GteNode) string {
  return ""
}

func (visitor BaseVisitor) VisitLtNode(node nodes.LtNode) string {
  return ""
}

func (visitor BaseVisitor) VisitLteNode(node nodes.LteNode) string {
  return ""
}

func (visitor BaseVisitor) VisitMatchesNode(node nodes.MatchesNode) string {
  return ""
}

func (visitor BaseVisitor) VisitOrNode(node nodes.OrNode) string {
  return ""
}

func (visitor BaseVisitor) VisitAttribute(node nodes.Attribute) string {
  return ""
}

func (visitor BaseVisitor) quotedAttribute(table, column string) string {
  if column == "*" {
    return fmt.Sprintf("%v.*", visitor.quote(table))
  }
  return fmt.Sprintf("%v.%v", visitor.quote(table), visitor.quote(column))
}

func (visitor BaseVisitor) quote(el string) string {
  return fmt.Sprintf(`"%v"`, el)
}
