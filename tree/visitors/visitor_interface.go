package visitors

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

type VisitorInterface interface {
  // Base methods.
  Accept(interface{}) string
  Visit(interface{}, VisitorInterface) string

  // Node visitors.
  VisitAssignment(*nodes.Assignment, VisitorInterface) string
  VisitEqual(*nodes.Equal, VisitorInterface) string
  VisitNotEqual(*nodes.NotEqual, VisitorInterface) string
  VisitGreaterThan(*nodes.GreaterThan, VisitorInterface) string
  VisitGreaterThanOrEqual(*nodes.GreaterThanOrEqual, VisitorInterface) string
  VisitLessThan(*nodes.LessThan, VisitorInterface) string
  VisitLessThanOrEqual(*nodes.LessThanOrEqual, VisitorInterface) string
  VisitLike(*nodes.Like, VisitorInterface) string
  VisitUnlike(*nodes.Unlike, VisitorInterface) string
  VisitOr(*nodes.Or, VisitorInterface) string
  VisitAnd(*nodes.And, VisitorInterface) string
  VisitRelation(*nodes.Relation, VisitorInterface) string
  VisitAttribute(*nodes.Attribute, VisitorInterface) string
  VisitGrouping(*nodes.Grouping, VisitorInterface) string
  VisitNot(*nodes.Not, VisitorInterface) string
  VisitLiteral(*nodes.Literal, VisitorInterface) string
  VisitInnerJoin(*nodes.InnerJoin, VisitorInterface) string
  VisitOuterJoin(*nodes.OuterJoin, VisitorInterface) string
  VisitOn(*nodes.On, VisitorInterface) string
  VisitUnqualifiedColumn(*nodes.UnqualifiedColumn, VisitorInterface) string
  VisitLimit(*nodes.Limit, VisitorInterface) string
  VisitOffset(*nodes.Offset, VisitorInterface) string
  VisitJoinSource(*nodes.JoinSource, VisitorInterface) string
  VisitSelectCore(*nodes.SelectCore, VisitorInterface) string
  VisitSelectStatement(*nodes.SelectStatement, VisitorInterface) string
  VisitValues(*nodes.Values, VisitorInterface) string
  VisitInsertStatement(*nodes.InsertStatement, VisitorInterface) string
  VisitUpdateStatement(*nodes.UpdateStatement, VisitorInterface) string
  VisitDeleteStatement(*nodes.DeleteStatement, VisitorInterface) string

  // SQL Functions.
  VisitCount(*nodes.Count, VisitorInterface) string
  VisitAverage(*nodes.Average, VisitorInterface) string
  VisitSum(*nodes.Sum, VisitorInterface) string
  VisitMaximum(*nodes.Maximum, VisitorInterface) string
  VisitMinimum(*nodes.Minimum, VisitorInterface) string

  // Base visitors.
  VisitString(interface{}) string
  VisitInteger(interface{}) string
  VisitFloat(interface{}) string
  VisitBool(interface{}) string

  // SQL Helpers.
  QuoteTableName(interface{}) string
  QuoteColumnName(interface{}) string
}
