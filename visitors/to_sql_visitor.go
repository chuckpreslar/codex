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
  case *nodes.MatchesNode:
    return visitor.visitMatchesNode(item.(*nodes.MatchesNode))
  case *nodes.OrNode:
    return visitor.visitOrNode(item.(*nodes.OrNode))
  case *nodes.SqlFunctionNode:
    return visitor.visitSqlFunctionNode(item.(*nodes.SqlFunctionNode))
  case *nodes.AttributeNode:
    return visitor.visitAttributeNode(item.(*nodes.AttributeNode))
  case *nodes.ReferenceNode:
    return visitor.visitReferenceNode(item.(*nodes.ReferenceNode))
  case *nodes.SelectCoreNode:
    return visitor.visitSelectCoreNode(item.(*nodes.SelectCoreNode))
  case string:
    return visitor.visitString(item.(string))
  case int:
    return visitor.visitInt(item.(int))
  case float64:
    return visitor.visitFloat64(item.(float64))
  default:
    panic("Unimplemented visitor method.")
  }
}

func (visitor *ToSqlVisitor) visitEqNode(eq *nodes.EqNode) string {
  if nil == eq.Right() {
    return fmt.Sprintf("%v IS NULL", visitor.visit(eq.Left()))
  }
  return fmt.Sprintf("%v = %v", visitor.visit(eq.Left()),
    utils.Tag(visitor.visit(eq.Right())))
}

func (visitor *ToSqlVisitor) visitNeqNode(neq *nodes.NeqNode) string {
  if nil == neq.Right() {
    return fmt.Sprintf("%v IS NOT NULL", visitor.visit(neq.Left()))
  }
  return fmt.Sprintf("%v != %v", visitor.visit(neq.Left()),
    utils.Tag(visitor.visit(neq.Right())))
}

func (visitor *ToSqlVisitor) visitGtNode(gt *nodes.GtNode) string {
  return fmt.Sprintf("%v > %v", visitor.visit(gt.Left()),
    utils.Tag(visitor.visit(gt.Right())))
}

func (visitor *ToSqlVisitor) visitGteNode(gte *nodes.GteNode) string {
  return fmt.Sprintf("%v >= %v", visitor.visit(gte.Left()),
    utils.Tag(visitor.visit(gte.Right())))
}

func (visitor *ToSqlVisitor) visitLtNode(lt *nodes.LtNode) string {
  return fmt.Sprintf("%v < %v", visitor.visit(lt.Left()),
    utils.Tag(visitor.visit(lt.Right())))
}

func (visitor *ToSqlVisitor) visitLteNode(lte *nodes.LteNode) string {
  return fmt.Sprintf("%v <= %v", visitor.visit(lte.Left()),
    utils.Tag(visitor.visit(lte.Right())))
}

func (visitor *ToSqlVisitor) visitMatchesNode(matches *nodes.MatchesNode) string {
  return fmt.Sprintf("%v LIKE %v", visitor.visit(matches.Left()),
    utils.Tag(visitor.visit(matches.Right())))
}

func (visitor *ToSqlVisitor) visitOrNode(or *nodes.OrNode) string {
  return fmt.Sprintf("%v OR %v", visitor.visit(or.Left()),
    visitor.visit(or.Right()))
}

func (visitor *ToSqlVisitor) visitSqlFunctionNode(function *nodes.SqlFunctionNode) string {
  return fmt.Sprintf("%v(%v)", visitor.functionName(function.FunctionName()),
    visitor.visit(function.Left()))
}

func (visitor *ToSqlVisitor) visitAttributeNode(attribute *nodes.AttributeNode) string {
  return fmt.Sprintf("%v.%v", visitor.visit(attribute.Left()),
    utils.Quote(visitor.visit(attribute.Right())))
}

func (visitor *ToSqlVisitor) visitReferenceNode(reference *nodes.ReferenceNode) string {
  return fmt.Sprintf("%v", utils.Quote(visitor.visit(reference.Left())))
}

func (visitor *ToSqlVisitor) visitSelectCoreNode(core *nodes.SelectCoreNode) (sql string) {
  if 0 < len(core.Projections()) {
    projections := make([]string, len(core.Projections()))
    for index, projection := range core.Projections() {
      projections[index] = visitor.visit(projection)
    }
    sql += strings.Join(projections, ",")
  } else {
    sql += visitor.literal(core.Reference())
  }
  sql += " FROM " + utils.Quote(core.Reference().Name())
  if 0 < len(core.Wheres()) {
    wheres := make([]string, len(core.Wheres()))
    for index, where := range core.Wheres() {
      wheres[index] = visitor.visit(where)
    }
    sql += fmt.Sprintf(" WHERE %s", strings.Join(wheres, ","))
  }
  return strings.Trim(sql, " ")
}

func (visitor *ToSqlVisitor) visitString(str string) string {
  return str
}

func (visitor *ToSqlVisitor) visitInt(i int) string {
  return fmt.Sprintf("%d", i)
}

func (visitor *ToSqlVisitor) visitFloat64(f float64) string {
  return strings.TrimRight(fmt.Sprintf("%f", f), "0")
}

func (visitor *ToSqlVisitor) functionName(function string) string {
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

func (visitor *ToSqlVisitor) literal(reference nodes.ReferenceInterface) string {
  return fmt.Sprintf(`"%v".*`, reference.Name())
}

func ToSql() *ToSqlVisitor {
  return &ToSqlVisitor{}
}
