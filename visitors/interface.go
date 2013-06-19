package visitors

import (
  "librarian/nodes"
)

type VisitorInterface interface {
  Accept(interface{}) string
  Visit(interface{}) string
  VisitComparisonNode(*nodes.ComparisonNode) string
  VisitEqualsNode(*nodes.EqualsNode) string
  VisitNotEqualsNode(*nodes.NotEqualsNode) string
  VisitGreaterThanNode(*nodes.GreaterThanNode) string
  VisitGreaterThanOrEqualsNode(*nodes.GreaterThanOrEqualsNode) string
  VisitLessThanNode(*nodes.LessThanNode) string
  VisitLessThanOrEqualsNode(*nodes.LessThanOrEqualsNode) string
  VisitMatchesNode(*nodes.MatchesNode) string
  VisitDoesNotMatchNode(*nodes.DoesNotMatchNode) string
  VisitAsNode(*nodes.AsNode) string
  VisitOrNode(*nodes.OrNode) string
  VisitOnNode(*nodes.OnNode) string
  VisitAndNode(*nodes.AndNode) string
  VisitAttributeNode(*nodes.AttributeNode) string
  VisitRelationNode(*nodes.RelationNode) string
  VisitJoinSourceNode(*nodes.JoinSourceNode) string
  VisitInnerJoinNode(*nodes.InnerJoinNode) string
  VisitLimitNode(*nodes.LimitNode) string
  VisitOffsetNode(*nodes.OffsetNode) string
  VisitSelectCoreNode(*nodes.SelectCoreNode) string
  VisitInt(int) string
  VisitString(string) string
  Quote(interface{}) string
  Tag(interface{}) string
}
