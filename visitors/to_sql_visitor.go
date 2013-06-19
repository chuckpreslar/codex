package visitors

import (
  "fmt"
  "librarian/nodes"
  "strings"
)

const (
  SELECT   = ` SELECT `
  FROM     = ` FROM `
  WHERE    = ` WHERE `
  SPACE    = ` `
  COMMA    = `, `
  GROUP_BY = ` GROUP BY `
  AND      = ` AND `
  DISINCT  = `DISTINCT`
  STAR     = `*`
)

type ToSqlVisitor struct{}

func (visitor *ToSqlVisitor) Accept(node interface{}) string {
  return visitor.Visit(node)
}

func (visitor *ToSqlVisitor) Visit(node interface{}) string {
  switch node.(type) {
  case *nodes.ComparisonNode:
    return visitor.VisitComparisonNode(node.(*nodes.ComparisonNode))
  case *nodes.EqualsNode:
    return visitor.VisitEqualsNode(node.(*nodes.EqualsNode))
  case *nodes.NotEqualsNode:
    return visitor.VisitNotEqualsNode(node.(*nodes.NotEqualsNode))
  case *nodes.GreaterThanNode:
    return visitor.VisitGreaterThanNode(node.(*nodes.GreaterThanNode))
  case *nodes.GreaterThanOrEqualsNode:
    return visitor.VisitGreaterThanOrEqualsNode(node.(*nodes.GreaterThanOrEqualsNode))
  case *nodes.LessThanNode:
    return visitor.VisitLessThanNode(node.(*nodes.LessThanNode))
  case *nodes.LessThanOrEqualsNode:
    return visitor.VisitLessThanOrEqualsNode(node.(*nodes.LessThanOrEqualsNode))
  case *nodes.MatchesNode:
    return visitor.VisitMatchesNode(node.(*nodes.MatchesNode))
  case *nodes.DoesNotMatchNode:
    return visitor.VisitDoesNotMatchNode(node.(*nodes.DoesNotMatchNode))
  case *nodes.AsNode:
    return visitor.VisitAsNode(node.(*nodes.AsNode))
  case *nodes.OrNode:
    return visitor.VisitOrNode(node.(*nodes.OrNode))
  case *nodes.OnNode:
    return visitor.VisitOnNode(node.(*nodes.OnNode))
  case *nodes.AndNode:
    return visitor.VisitAndNode(node.(*nodes.AndNode))
  case *nodes.AttributeNode:
    return visitor.VisitAttributeNode(node.(*nodes.AttributeNode))
  case *nodes.RelationNode:
    return visitor.VisitRelationNode(node.(*nodes.RelationNode))
  case *nodes.JoinSourceNode:
    return visitor.VisitJoinSourceNode(node.(*nodes.JoinSourceNode))
  case *nodes.InnerJoinNode:
    return visitor.VisitInnerJoinNode(node.(*nodes.InnerJoinNode))
  case *nodes.LimitNode:
    return visitor.VisitLimitNode(node.(*nodes.LimitNode))
  case *nodes.OffsetNode:
    return visitor.VisitOffsetNode(node.(*nodes.OffsetNode))
  case *nodes.SelectCoreNode:
    return visitor.VisitSelectCoreNode(node.(*nodes.SelectCoreNode))
  case *nodes.SelectStatementNode:
    return visitor.VisitSelectStatementNode(node.(*nodes.SelectStatementNode))
  case int:
    return visitor.VisitInt(node.(int))
  case string:
    return visitor.VisitString(node.(string))
  case bool:
    return visitor.VisitBool(node.(bool))
  default:
    panic(fmt.Sprintf("Unkown Node type: %T", node))
  }
}

func (visitor *ToSqlVisitor) VisitComparisonNode(comparison *nodes.ComparisonNode) string {
  return visitor.Visit(comparison.Left)
}

func (visitor *ToSqlVisitor) VisitEqualsNode(eq *nodes.EqualsNode) string {
  if nil == eq.Right {
    return fmt.Sprintf("%v IS NULL", visitor.Visit(eq.Left))
  }
  return fmt.Sprintf("%v = %v", visitor.Visit(eq.Left),
    visitor.Visit(eq.Right))
}

func (visitor *ToSqlVisitor) VisitNotEqualsNode(neq *nodes.NotEqualsNode) string {
  if nil == neq.Right {
    return fmt.Sprintf("%v IS NOT NULL", visitor.Visit(neq.Left))
  }
  return fmt.Sprintf("%v != %v", visitor.Visit(neq.Left),
    visitor.Visit(neq.Right))
}

func (visitor *ToSqlVisitor) VisitGreaterThanNode(gt *nodes.GreaterThanNode) string {
  return fmt.Sprintf("%v > %v", visitor.Visit(gt.Left),
    visitor.Visit(gt.Right))
}

func (visitor *ToSqlVisitor) VisitGreaterThanOrEqualsNode(gte *nodes.GreaterThanOrEqualsNode) string {
  return fmt.Sprintf("%v >= %v", visitor.Visit(gte.Left),
    visitor.Visit(gte.Right))
}
func (visitor *ToSqlVisitor) VisitLessThanNode(lt *nodes.LessThanNode) string {
  return fmt.Sprintf("%v < %v", visitor.Visit(lt.Left),
    visitor.Visit(lt.Right))
}

func (visitor *ToSqlVisitor) VisitLessThanOrEqualsNode(lte *nodes.LessThanOrEqualsNode) string {
  return fmt.Sprintf("%v <= %v", visitor.Visit(lte.Left),
    visitor.Visit(lte.Right))
}

func (visitor *ToSqlVisitor) VisitMatchesNode(matches *nodes.MatchesNode) string {
  return fmt.Sprintf("%v LIKE %v", visitor.Visit(matches.Left),
    visitor.Visit(matches.Right))
}

func (visitor *ToSqlVisitor) VisitDoesNotMatchNode(dnm *nodes.DoesNotMatchNode) string {
  return fmt.Sprintf("%v NOT LIKE %v", visitor.Visit(dnm.Left),
    visitor.Visit(dnm.Right))
}

func (visitor *ToSqlVisitor) VisitAsNode(as *nodes.AsNode) string {
  return fmt.Sprintf("%v AS %v", visitor.Visit(as.Left),
    visitor.Visit(as.Right))
}

func (visitor *ToSqlVisitor) VisitOrNode(or *nodes.OrNode) string {
  return fmt.Sprintf("%v OR %v", visitor.Visit(or.Left),
    visitor.Visit(or.Right))
}

func (visitor *ToSqlVisitor) VisitAndNode(and *nodes.AndNode) string {
  return fmt.Sprintf("%v AND %v", visitor.Visit(and.Left),
    visitor.Visit(and.Right))
}

func (visitor *ToSqlVisitor) VisitOnNode(on *nodes.OnNode) string {
  return fmt.Sprintf("%v ON %v", visitor.Visit(on.Left),
    visitor.Visit(on.Right))
}

func (visitor *ToSqlVisitor) VisitAttributeNode(attribute *nodes.AttributeNode) string {
  var relation string
  if 0 < len(attribute.Relation.Alias) {
    relation = attribute.Relation.Alias
  } else {
    relation = attribute.Relation.Name
  }
  return fmt.Sprintf("%v.%v", visitor.Quote(relation), visitor.Quote(attribute.Name))
}

func (visitor *ToSqlVisitor) VisitRelationNode(relation *nodes.RelationNode) string {
  var name string
  if 0 < len(relation.Alias) {
    name = relation.Alias
  } else {
    name = relation.Name
  }
  return visitor.Quote(name)
}

func (visitor *ToSqlVisitor) VisitJoinSourceNode(source *nodes.JoinSourceNode) string {
  str := fmt.Sprintf("%v", visitor.Visit(source.Left))
  for _, join := range source.Right {
    str = fmt.Sprintf("%v %v ", str, visitor.Visit(join))
  }
  return visitor.Trim(str)
}

func (visitor *ToSqlVisitor) VisitInnerJoinNode(join *nodes.InnerJoinNode) string {
  return fmt.Sprintf("INNER JOIN %v", visitor.Visit(join.Left))
}

func (visitor *ToSqlVisitor) VisitLimitNode(limit *nodes.LimitNode) string {
  return fmt.Sprintf("LIMIT %v", visitor.Visit(limit.Left))
}

func (visitor *ToSqlVisitor) VisitOffsetNode(offset *nodes.OffsetNode) string {
  return fmt.Sprintf("OFFSET %v", visitor.Visit(offset.Left))
}

func (visitor *ToSqlVisitor) VisitSelectCoreNode(core *nodes.SelectCoreNode) string {
  str := fmt.Sprintf("%v", SELECT)

  if length := len(core.Projections) - 1; 0 <= length {
    for index, projection := range core.Projections {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(projection))

      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
  } else {
    str = fmt.Sprintf("%v%v.%v", str, visitor.Visit(core.Relation), "*")
  }

  str = fmt.Sprintf("%v%v%v", str, FROM, visitor.Visit(core.Source))

  if length := len(core.Wheres) - 1; 0 <= length {
    str = fmt.Sprintf("%v%v", str, WHERE)
    for index, where := range core.Wheres {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(where))
      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
  }

  return visitor.Trim(str)
}

func (visitor *ToSqlVisitor) VisitSelectStatementNode(stmt *nodes.SelectStatementNode) string {
  str := ""
  for _, core := range stmt.Cores {
    str = fmt.Sprintf("%v%v", str, visitor.Visit(core))
  }

  if nil != stmt.Limit {
    str = fmt.Sprintf("%v %v", str, visitor.Visit(stmt.Limit))
  }

  if nil != stmt.Offset {
    str = fmt.Sprintf("%v %v", str, visitor.Visit(stmt.Offset))
  }

  return str
}

func (visitor *ToSqlVisitor) VisitInt(integer int) string {
  return fmt.Sprintf("%d", integer)
}

func (visitor *ToSqlVisitor) VisitString(str string) string {
  return fmt.Sprintf("%s", visitor.Tag(str))
}

func (visitor *ToSqlVisitor) VisitBool(boolean bool) string {
  return fmt.Sprintf("%s", visitor.Tag(boolean))
}

// Utility functions.

func (visitor *ToSqlVisitor) Quote(value interface{}) string {
  return fmt.Sprintf(`"%v"`, value)
}

func (visitor *ToSqlVisitor) Tag(value interface{}) string {
  return fmt.Sprintf(`'%v'`, value)
}

func (visitor *ToSqlVisitor) Trim(value interface{}) string {
  return strings.Trim(strings.TrimRight(fmt.Sprintf("%v", value), "0"), " ")
}

func ToSql() *ToSqlVisitor {
  return &ToSqlVisitor{}
}
