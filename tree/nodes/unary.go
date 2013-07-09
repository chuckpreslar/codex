package nodes

// Unary node struct
type UnaryNode struct {
  Expr interface{} // Single leaf for the Unary node.
}

type LiteralNode UnaryNode           // Literal node is a Unary node struct
type OnNode UnaryNode                // On node is a Unary node struct
type LimitNode UnaryNode             // Limit node is a Unary node struct
type OffsetNode UnaryNode            // Offset node is a Unary node struct
type GroupNode UnaryNode             // Group node is a Unary node struct
type HavingNode UnaryNode            // Having node is a Unary node struct
type UnqualifiedColumnNode UnaryNode // UnqualifiedColumn node is a Unary node struct
type ColumnNode UnaryNode            // Column node is a Unary node struct
type StarNode UnaryNode              // Star node is a Unary node struct

func Literal(expr interface{}) *LiteralNode {
  literal := new(LiteralNode)
  literal.Expr = expr
  return literal
}

func On(expr interface{}) *OnNode {
  on := new(OnNode)
  on.Expr = expr
  return on
}

func Limit(expr interface{}) *LimitNode {
  limit := new(LimitNode)
  limit.Expr = expr
  return limit
}

func Offset(expr interface{}) *OffsetNode {
  offset := new(OffsetNode)
  offset.Expr = expr
  return offset
}

func Group(expr interface{}) *GroupNode {
  group := new(GroupNode)
  group.Expr = expr
  return group
}

func Having(expr interface{}) *HavingNode {
  having := new(HavingNode)
  having.Expr = expr
  return having
}

func UnqualifiedColumn(expr interface{}) *UnqualifiedColumnNode {
  column := new(UnqualifiedColumnNode)
  column.Expr = expr
  return column
}

func Column(expr interface{}) *ColumnNode {
  column := new(ColumnNode)
  column.Expr = expr
  return column
}

func Star() *StarNode {
  return new(StarNode)
}
