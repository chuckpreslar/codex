package visitors

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

type VisitorInterface interface {
  // Base methods.
  Accept(interface{}) string
  Visit(interface{}, VisitorInterface) string

  // Node visitors.
  VisitAssignment(*nodes.AssignmentNode, VisitorInterface) string
  VisitEqual(*nodes.EqualNode, VisitorInterface) string
  VisitNotEqual(*nodes.NotEqualNode, VisitorInterface) string
  VisitGreaterThan(*nodes.GreaterThanNode, VisitorInterface) string
  VisitGreaterThanOrEqual(*nodes.GreaterThanOrEqualNode, VisitorInterface) string
  VisitLessThan(*nodes.LessThanNode, VisitorInterface) string
  VisitLessThanOrEqual(*nodes.LessThanOrEqualNode, VisitorInterface) string
  VisitLike(*nodes.LikeNode, VisitorInterface) string
  VisitUnlike(*nodes.UnlikeNode, VisitorInterface) string
  VisitOr(*nodes.OrNode, VisitorInterface) string
  VisitAnd(*nodes.AndNode, VisitorInterface) string
  VisitRelation(*nodes.RelationNode, VisitorInterface) string
  VisitAttribute(*nodes.AttributeNode, VisitorInterface) string
  VisitGrouping(*nodes.GroupingNode, VisitorInterface) string
  VisitNot(*nodes.NotNode, VisitorInterface) string
  VisitLiteral(*nodes.LiteralNode, VisitorInterface) string
  VisitInnerJoin(*nodes.InnerJoinNode, VisitorInterface) string
  VisitOuterJoin(*nodes.OuterJoinNode, VisitorInterface) string
  VisitOn(*nodes.OnNode, VisitorInterface) string
  VisitUnqualifiedColumn(*nodes.UnqualifiedColumnNode, VisitorInterface) string
  VisitLimit(*nodes.LimitNode, VisitorInterface) string
  VisitOffset(*nodes.OffsetNode, VisitorInterface) string
  VisitJoinSource(*nodes.JoinSourceNode, VisitorInterface) string
  VisitSelectCore(*nodes.SelectCoreNode, VisitorInterface) string
  VisitSelectStatement(*nodes.SelectStatementNode, VisitorInterface) string
  VisitValues(*nodes.ValuesNode, VisitorInterface) string
  VisitInsertStatement(*nodes.InsertStatementNode, VisitorInterface) string
  VisitUpdateStatement(*nodes.UpdateStatementNode, VisitorInterface) string
  VisitDeleteStatement(*nodes.DeleteStatementNode, VisitorInterface) string

  // SQL Functions.
  VisitCount(*nodes.CountNode, VisitorInterface) string
  VisitAverage(*nodes.AverageNode, VisitorInterface) string
  VisitSum(*nodes.SumNode, VisitorInterface) string
  VisitMaximum(*nodes.MaximumNode, VisitorInterface) string
  VisitMinimum(*nodes.MinimumNode, VisitorInterface) string

  // Base visitors.
  VisitString(interface{}) string
  VisitInteger(interface{}) string
  VisitFloat(interface{}) string
  VisitBool(interface{}) string

  // SQL Helpers.
  QuoteTableName(interface{}) string
  QuoteColumnName(interface{}) string
}
