package visitors

import (
  "librarian/nodes"
)

type VisitComparisonInterface interface {
  Accept(nodes.NodeInterface) string
  visit(interface{}) string
  visitEqNode(*nodes.EqNode) string
  visitNeqNode(*nodes.NeqNode) string
  visitGtNode(*nodes.GtNode) string
  visitGteNode(*nodes.GteNode) string
  visitLtNode(*nodes.LtNode) string
  visitLteNode(*nodes.LteNode) string
  visitMatchesNode(*nodes.MatchesNode) string
  visitOrNode(*nodes.OrNode) string
  visitSqlFunctionNode(*nodes.SqlFunctionNode) string
  visitAttributeNode(*nodes.AttributeNode) string
  visitSelectCoreNode(*nodes.SelectCoreNode) string
  visitString(string) string
  visitInt(int) string
  visitFloat64(float64) string
  functionName(string) string
}
