package visitors

import (
  "codex/tree/nodes"
)

type VisitorInterface interface {
  Accept(interface{}) string
  Visit(interface{}) string
  VisitAssignment(*nodes.Assignment) string
  VisitEqual(*nodes.Equal) string
  VisitNotEqual(*nodes.NotEqual) string
  VisitGreaterThan(*nodes.GreaterThan) string
  VisitGreaterThanOrEqual(*nodes.GreaterThanOrEqual) string
  VisitLessThan(*nodes.LessThan) string
  VisitLessThanOrEqual(*nodes.LessThanOrEqual) string
  VisitLike(*nodes.Like) string
  VisitUnlike(*nodes.Unlike) string
  VisitOr(*nodes.Or) string
  VisitAnd(*nodes.And) string
  VisitRelation(*nodes.Relation) string
  VisitAttribute(*nodes.Attribute) string
  VisitGrouping(*nodes.Grouping) string
  VisitNot(*nodes.Not) string
  VisitLiteral(*nodes.Literal) string
  VisitInnerJoin(*nodes.InnerJoin) string
  VisitOuterJoin(*nodes.OuterJoin) string
  VisitOn(*nodes.On) string
  VisitUnqualifiedColumn(*nodes.UnqualifiedColumn) string
  VisitLimit(*nodes.Limit) string
  VisitOffset(*nodes.Offset) string
  VisitJoinSource(*nodes.JoinSource) string
  VisitSelectCore(*nodes.SelectCore) string
  VisitSelectStatement(*nodes.SelectStatement) string
  VisitValues(*nodes.Values) string
  VisitInsertStatement(*nodes.InsertStatement) string
  VisitUpdateStatement(*nodes.UpdateStatement) string
  VisitDeleteStatement(*nodes.DeleteStatement) string
  VisitString(interface{}) string
  VisitInteger(interface{}) string
  VisitFloat(interface{}) string
  VisitBool(interface{}) string
}
