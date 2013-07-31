package visitors

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

type VisitorInterface interface {
  // Base methods.
  Accept(interface{}) (string, error)
  Visit(interface{}, VisitorInterface) string

  // Unary node visitors.
  VisitGrouping(*nodes.GroupingNode, VisitorInterface) string
  VisitNot(*nodes.NotNode, VisitorInterface) string
  VisitLiteral(*nodes.LiteralNode, VisitorInterface) string
  VisitOn(*nodes.OnNode, VisitorInterface) string
  VisitColumn(*nodes.ColumnNode, VisitorInterface) string
  VisitStar(*nodes.StarNode, VisitorInterface) string
  VisitBinding(*nodes.BindingNode, VisitorInterface) string
  VisitUnqualifiedColumn(*nodes.UnqualifiedColumnNode, VisitorInterface) string
  VisitLimit(*nodes.LimitNode, VisitorInterface) string
  VisitOffset(*nodes.OffsetNode, VisitorInterface) string
  VisitHaving(*nodes.HavingNode, VisitorInterface) string
  VisitAscending(*nodes.AscendingNode, VisitorInterface) string
  VisitDescending(*nodes.DescendingNode, VisitorInterface) string

  // Binary node visitors.
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
  VisitInnerJoin(*nodes.InnerJoinNode, VisitorInterface) string
  VisitOuterJoin(*nodes.OuterJoinNode, VisitorInterface) string
  VisitJoinSource(*nodes.JoinSourceNode, VisitorInterface) string
  VisitValues(*nodes.ValuesNode, VisitorInterface) string
  VisitUnion(*nodes.UnionNode, VisitorInterface) string
  VisitExcept(*nodes.ExceptNode, VisitorInterface) string
  VisitIntersect(*nodes.IntersectNode, VisitorInterface) string

  // Nary node visitors.
  VisitSelectCore(*nodes.SelectCoreNode, VisitorInterface) string
  VisitSelectStatement(*nodes.SelectStatementNode, VisitorInterface) string
  VisitInsertStatement(*nodes.InsertStatementNode, VisitorInterface) string
  VisitUpdateStatement(*nodes.UpdateStatementNode, VisitorInterface) string
  VisitDeleteStatement(*nodes.DeleteStatementNode, VisitorInterface) string

  // Function node visitors.
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

  // Helpers.
  QuoteTableName(interface{}) string
  QuoteColumnName(interface{}) string
}
