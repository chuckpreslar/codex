package visitors

import (
  "fmt"
  "librarian/nodes"
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

func (visitor *ToSqlVisitor) Visit(item interface{}) string {
  switch item.(type) {
  case *nodes.ComparisonNode:
    return visitor.VisitComparisonNode(item.(*nodes.ComparisonNode))
  case *nodes.EqualsNode:
    return visitor.VisitEqualsNode(item.(*nodes.EqualsNode))
  case *nodes.NotEqualsNode:
    return visitor.VisitNotEqualsNode(item.(*nodes.NotEqualsNode))
  case *nodes.GreaterThanNode:
    return visitor.VisitGreaterThanNode(item.(*nodes.GreaterThanNode))
  case *nodes.GreaterThanOrEqualsNode:
    return visitor.VisitGreaterThanOrEqualsNode(item.(*nodes.GreaterThanOrEqualsNode))
  case *nodes.LessThanNode:
    return visitor.VisitLessThanNode(item.(*nodes.LessThanNode))
  case *nodes.LessThanOrEqualsNode:
    return visitor.VisitLessThanOrEqualsNode(item.(*nodes.LessThanOrEqualsNode))
  case *nodes.MatchesNode:
    return visitor.VisitMatchesNode(item.(*nodes.MatchesNode))
  case *nodes.DoesNotMatchNode:
    return visitor.VisitDoesNotMatchNode(item.(*nodes.DoesNotMatchNode))
  case *nodes.AsNode:
    return visitor.VisitAsNode(item.(*nodes.AsNode))
  case *nodes.OrNode:
    return visitor.VisitOrNode(item.(*nodes.OrNode))
  case *nodes.AndNode:
    return visitor.VisitAndNode(item.(*nodes.AndNode))
  case int:
    return visitor.VisitInt(item.(int))
  default:
    panic("Unkown Node type.")
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

func (visitor *ToSqlVisitor) VisitInt(integer int) string {
  return fmt.Sprintf("%d", integer)
}

func ToSql() *ToSqlVisitor {
  return &ToSqlVisitor{}
}
