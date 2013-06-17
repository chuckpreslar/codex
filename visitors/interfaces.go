package visitors

import (
  "librarian/nodes"
)

type VisitorInterface interface {
  Accept(nodes.NodeInterface) string
  Visit(interface{}) string
  VisitEqNode(*nodes.EqNode) string
  VisitNeqNode(*nodes.NeqNode) string
  VisitGtNode(*nodes.GtNode) string
  VisitGteNode(*nodes.GteNode) string
  VisitLtNode(*nodes.LtNode) string
  VisitLteNode(*nodes.LteNode) string
  VisitMatchesNode(*nodes.MatchesNode) string
  VisitOrNode(*nodes.OrNode) string
  VisitSqlFunctionNode(*nodes.SqlFunctionNode) string
  VisitAttributeNode(*nodes.AttributeNode) string
  VisitSelectCoreNode(*nodes.SelectCoreNode) string
  VisitString(string) string
  VisitInt(int) string
  VisitFloat64(float64) string
  FunctionName(string) string
}
