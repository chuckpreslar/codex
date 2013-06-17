package visitors

import (
  "fmt"
  "librarian/nodes"
  "librarian/visitors/utils"
  "strings"
)

type ToSqlVisitor struct {
  ast nodes.NodeInterface
}

func (visitor *ToSqlVisitor) Accept(ast nodes.NodeInterface) string {
  visitor.ast = ast
  return visitor.Visit(ast)
}

func (visitor *ToSqlVisitor) Visit(item interface{}) string {
  switch item.(type) {
  case *nodes.EqNode:
    return visitor.VisitEqNode(item.(*nodes.EqNode))
  case *nodes.NeqNode:
    return visitor.VisitNeqNode(item.(*nodes.NeqNode))
  case *nodes.GtNode:
    return visitor.VisitGtNode(item.(*nodes.GtNode))
  case *nodes.GteNode:
    return visitor.VisitGteNode(item.(*nodes.GteNode))
  case *nodes.LtNode:
    return visitor.VisitLtNode(item.(*nodes.LtNode))
  case *nodes.LteNode:
    return visitor.VisitLteNode(item.(*nodes.LteNode))
  case *nodes.MatchesNode:
    return visitor.VisitMatchesNode(item.(*nodes.MatchesNode))
  case *nodes.OrNode:
    return visitor.VisitOrNode(item.(*nodes.OrNode))
  case *nodes.SqlFunctionNode:
    return visitor.VisitSqlFunctionNode(item.(*nodes.SqlFunctionNode))
  case *nodes.AttributeNode:
    return visitor.VisitAttributeNode(item.(*nodes.AttributeNode))
  case *nodes.RelationNode:
    return visitor.VisitRelationNode(item.(*nodes.RelationNode))
  case *nodes.SelectCoreNode:
    return visitor.VisitSelectCoreNode(item.(*nodes.SelectCoreNode))
  case string:
    return visitor.VisitString(item.(string))
  case int:
    return visitor.VisitInt(item.(int))
  case float64:
    return visitor.VisitFloat64(item.(float64))
  default:
    panic("Unimplemented visitor method.")
  }
}

func (visitor *ToSqlVisitor) VisitEqNode(eq *nodes.EqNode) string {
  if nil == eq.Right {
    return fmt.Sprintf("%v IS NULL", visitor.Visit(eq.Left))
  }
  return fmt.Sprintf("%v = %v", visitor.Visit(eq.Left),
    utils.Tag(visitor.Visit(eq.Right)))
}

func (visitor *ToSqlVisitor) VisitNeqNode(neq *nodes.NeqNode) string {
  if nil == neq.Right {
    return fmt.Sprintf("%v IS NOT NULL", visitor.Visit(neq.Left))
  }
  return fmt.Sprintf("%v != %v", visitor.Visit(neq.Left),
    utils.Tag(visitor.Visit(neq.Right)))
}

func (visitor *ToSqlVisitor) VisitGtNode(gt *nodes.GtNode) string {
  return fmt.Sprintf("%v > %v", visitor.Visit(gt.Left),
    utils.Tag(visitor.Visit(gt.Right)))
}

func (visitor *ToSqlVisitor) VisitGteNode(gte *nodes.GteNode) string {
  return fmt.Sprintf("%v >= %v", visitor.Visit(gte.Left),
    utils.Tag(visitor.Visit(gte.Right)))
}

func (visitor *ToSqlVisitor) VisitLtNode(lt *nodes.LtNode) string {
  return fmt.Sprintf("%v < %v", visitor.Visit(lt.Left),
    utils.Tag(visitor.Visit(lt.Right)))
}

func (visitor *ToSqlVisitor) VisitLteNode(lte *nodes.LteNode) string {
  return fmt.Sprintf("%v <= %v", visitor.Visit(lte.Left),
    utils.Tag(visitor.Visit(lte.Right)))
}

func (visitor *ToSqlVisitor) VisitMatchesNode(matches *nodes.MatchesNode) string {
  return fmt.Sprintf("%v LIKE %v", visitor.Visit(matches.Left),
    utils.Tag(visitor.Visit(matches.Right)))
}

func (visitor *ToSqlVisitor) VisitOrNode(or *nodes.OrNode) string {
  return fmt.Sprintf("%v OR %v", visitor.Visit(or.Left),
    visitor.Visit(or.Right))
}

func (visitor *ToSqlVisitor) VisitSqlFunctionNode(function *nodes.SqlFunctionNode) string {
  return fmt.Sprintf("%v(%v)", visitor.FunctionName(function.FunctionName()),
    visitor.Visit(function.Left))
}

func (visitor *ToSqlVisitor) VisitAttributeNode(attribute *nodes.AttributeNode) string {
  return fmt.Sprintf("%v.%v", visitor.Visit(attribute.Left),
    utils.Quote(visitor.Visit(attribute.Right)))
}

func (visitor *ToSqlVisitor) VisitRelationNode(relation *nodes.RelationNode) string {
  return fmt.Sprintf("%v", utils.Quote(visitor.Visit(relation.Left)))
}

func (visitor *ToSqlVisitor) VisitSelectCoreNode(core *nodes.SelectCoreNode) (sql string) {
  return strings.Trim(sql, " ")
}

func (visitor *ToSqlVisitor) VisitString(str string) string {
  return str
}

func (visitor *ToSqlVisitor) VisitInt(i int) string {
  return fmt.Sprintf("%d", i)
}

func (visitor *ToSqlVisitor) VisitFloat64(f float64) string {
  return strings.TrimRight(fmt.Sprintf("%f", f), "0")
}

func (visitor *ToSqlVisitor) FunctionName(function string) string {
  switch function = strings.ToLower(function); function {
  case "maximum":
    return "MAX"
  case "minimum":
    return "MIN"
  case "average":
    return "AVG"
  case "count":
    return "COUNT"
  case "first":
    return "FIRST"
  case "last":
    return "LAST"
  case "sum":
    return "SUM"
  case "upper":
    return "UPPER"
  case "lower":
    return "LOWER"
  case "mid":
    return "MID"
  case "length":
    return "LEN"
  case "round":
    return "ROUND"
  default:
    panic("Unkown SQL Function.")
  }
}

func (visitor *ToSqlVisitor) literal(relation nodes.RelationInterface) string {
  return fmt.Sprintf(`"%v".*`, relation.Name())
}

func ToSql() *ToSqlVisitor {
  return &ToSqlVisitor{}
}
