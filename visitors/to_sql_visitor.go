package visitors

import (
  "fmt"
  "librarian/nodes"
  "strings"
)

const (
  WHERE    = ` WHERE `
  SPACE    = ` `
  COMMA    = `, `
  GROUP_BY = ` GROUP BY `
  AND      = ` AND `
  DISINCT  = `DISTINCT`
  STAR     = `*`
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
  case *nodes.SelectStatementNode:
    return visitor.VisitSelectStatementNode(item.(*nodes.SelectStatementNode))
  case string:
    return visitor.VisitString(item.(string))
  case int:
    return visitor.VisitInt(item.(int))
  case float64:
    return visitor.VisitFloat64(item.(float64))
  case bool:
    return visitor.VisitBool(item.(bool))
  default:
    panic("Unimplemented visitor method.")
  }
}

func (visitor *ToSqlVisitor) VisitEqNode(eq *nodes.EqNode) string {
  if nil == eq.Right {
    return fmt.Sprintf("%v IS NULL", visitor.Visit(eq.Left))
  }
  return fmt.Sprintf("%v = %v", visitor.Visit(eq.Left),
    visitor.Visit(eq.Right))
}

func (visitor *ToSqlVisitor) VisitNeqNode(neq *nodes.NeqNode) string {
  if nil == neq.Right {
    return fmt.Sprintf("%v IS NOT NULL", visitor.Visit(neq.Left))
  }
  return fmt.Sprintf("%v != %v", visitor.Visit(neq.Left),
    visitor.Visit(neq.Right))
}

func (visitor *ToSqlVisitor) VisitGtNode(gt *nodes.GtNode) string {
  return fmt.Sprintf("%v > %v", visitor.Visit(gt.Left),
    visitor.Visit(gt.Right))
}

func (visitor *ToSqlVisitor) VisitGteNode(gte *nodes.GteNode) string {
  return fmt.Sprintf("%v >= %v", visitor.Visit(gte.Left),
    visitor.Visit(gte.Right))
}

func (visitor *ToSqlVisitor) VisitLtNode(lt *nodes.LtNode) string {
  return fmt.Sprintf("%v < %v", visitor.Visit(lt.Left),
    visitor.Visit(lt.Right))
}

func (visitor *ToSqlVisitor) VisitLteNode(lte *nodes.LteNode) string {
  return fmt.Sprintf("%v <= %v", visitor.Visit(lte.Left),
    visitor.Visit(lte.Right))
}

func (visitor *ToSqlVisitor) VisitMatchesNode(matches *nodes.MatchesNode) string {
  return fmt.Sprintf("%v LIKE %v", visitor.Visit(matches.Left),
    visitor.Visit(matches.Right))
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
    quote(attribute.Right))
}

func (visitor *ToSqlVisitor) VisitRelationNode(relation *nodes.RelationNode) string {
  return fmt.Sprintf("%v", quote(relation.Left))
}

func (visitor *ToSqlVisitor) VisitSelectCoreNode(core *nodes.SelectCoreNode) string {
  str := "SELECT "

  if 0 == len(core.Projections) {
    str = fmt.Sprintf("%v%v.%v", str, visitor.Visit(core.Relation), STAR)
  } else {
    for index, projection := range core.Projections {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(projection))
      if index+1 != len(core.Projections) {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
  }

  str += fmt.Sprintf(" FROM %v", visitor.Visit(core.Relation))

  if 0 != len(core.Wheres) {
    str = fmt.Sprintf("%v%v", str, WHERE)
    for index, where := range core.Wheres {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(where))
      if index+1 != len(core.Wheres) {
        str = fmt.Sprintf("%v%v", str, AND)
      }
    }
  }

  return strings.Trim(str, " ")
}

func (visitor *ToSqlVisitor) VisitSelectStatementNode(stmt *nodes.SelectStatementNode) string {
  return visitor.Visit(stmt.CoreAtIndex(0))
}

func (visitor *ToSqlVisitor) VisitString(str string) string {
  return tag(str)
}

func (visitor *ToSqlVisitor) VisitInt(i int) string {
  return tag(i)
}

func (visitor *ToSqlVisitor) VisitFloat64(f float64) string {
  return strings.TrimRight(tag(f), "0")
}

func (visitor *ToSqlVisitor) VisitBool(b bool) string {
  return tag(b)
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

func quote(value interface{}) string {
  switch value.(type) {
  case string:
    return fmt.Sprintf(`"%v"`, value)
  default:
    return fmt.Sprintf(`%v`, value)
  }
}

func ToSql() *ToSqlVisitor {
  return &ToSqlVisitor{}
}
